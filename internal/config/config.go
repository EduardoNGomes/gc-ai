package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eduardongomes/gcai/errs"
	linereader "github.com/eduardongomes/gcai/internal/line-reader"
	"github.com/eduardongomes/gcai/internal/prompt"
	"github.com/eduardongomes/gcai/internal/providers"
)

type Config struct {
	open_ai_key         string
	gemini_key          string
	configPath          string
	allow_edit          bool
	agent               providers.Provider
	promptType          prompt.PromptType
	prompt              prompt.Prompt
	menuPromptSelector  MenuPromptSelector
	customPromptFactory CustomPromptFactory
}

type RewriteConfigOptions struct {
	PromptType *prompt.PromptType
	Prompt     *prompt.Prompt
	AllowEdit  *bool
	Agent      *providers.Provider
}

type envStruct struct {
	GeminiKey    string                  `json:"geminiKey"`
	OpenAIKey    string                  `json:"openAIAKey"`
	AllowEdit    bool                    `json:"allowEdit"`
	Agent        providers.Provider      `json:"agent"`
	PromptType   prompt.PromptType       `json:"promptType"`
	CustomPrompt prompt.CustomPromptJSON `json:"customPrompt"`
}

var e envStruct

type ConfigMethods interface {
	LoadEnvs(configPath string) error
	IsEmpty() bool
	Config(io.Reader, providers.AgentOptions, io.Writer) error

	GetGeminiKey() string
	GetOpenAIKey() string

	GetAllowEdit() bool
	setAllowEdit(v bool)

	GetAgent() providers.Provider
	setAgent(providers.Provider)

	GetPromptType() prompt.PromptType
	setPromptType(prompt.PromptType)

	GetPromptString() string
	GetPrompt() prompt.Prompt
	setPrompt(prompt.Prompt)

	RewriteConfig(RewriteConfigOptions) error
}

func NewConfig() *Config {
	return &Config{
		menuPromptSelector:  ProdMenuSelector{},
		customPromptFactory: ProdCustomPromptFactory{},
	}
}

func (c *Config) IsEmpty() bool {
	gemini := c.GetGeminiKey()
	openAI := c.GetOpenAIKey()

	if len(gemini) == 0 && len(openAI) == 0 {
		return true
	}

	return false
}

func (c *Config) GetAllowEdit() bool {
	return c.allow_edit
}

func (c *Config) setAllowEdit(v bool) {
	c.allow_edit = v
}

func (c *Config) GetOpenAIKey() string {
	return c.open_ai_key
}

func (c *Config) setOpenAIKey(v string) {
	c.open_ai_key = v
}

func (c *Config) GetGeminiKey() string {
	return c.gemini_key
}

func (c *Config) setGeminiKey(v string) {
	c.gemini_key = v
}

func (c *Config) setConfigPath(v string) {
	c.configPath = v
}

func (c *Config) GetAgent() providers.Provider {
	return c.agent
}

func (c *Config) setAgent(v providers.Provider) {
	c.agent = v
}

func (c *Config) GetPromptString() string {
	return prompt.ConvertToPromptString(c.prompt)
}

func (c *Config) GetPrompt() prompt.Prompt {
	return c.prompt
}

func (c *Config) setPrompt(p prompt.Prompt) {
	c.prompt = p
}

func (c *Config) GetPromptType() prompt.PromptType {
	return c.promptType
}

func (c *Config) setPromptType(v prompt.PromptType) {
	c.promptType = v
}

func (c *Config) RewriteConfig(v RewriteConfigOptions) error {

	if v.AllowEdit != nil {
		c.setAllowEdit(*v.AllowEdit)
	}

	if v.Agent != nil {
		c.setAgent(*v.Agent)
	}

	if v.PromptType != nil {
		c.setPromptType(*v.PromptType)
	}

	if v.Prompt != nil {
		c.setPrompt(*v.Prompt)
	}

	fileConfig, err := os.OpenFile(c.configPath, os.O_RDWR, 0)

	if err != nil {
		return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
	}

	defer fileConfig.Close()

	cfg := configDTO(fileConfig, c)
	if err = writeConfig(cfg); err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadEnvs(configPath string) error {

	f, err := os.ReadFile(configPath)

	if err != nil {
		return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
	}

	err = json.Unmarshal(f, &e)

	if err := json.Unmarshal(f, &e); err != nil {
		fmt.Println("err", err)
		fileConfig, err := os.OpenFile(configPath, os.O_RDWR, 0)

		if err != nil {
			return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
		}

		defer fileConfig.Close()

		cfg := configDTO(fileConfig, c)
		writeConfig(cfg)
	}

	c.setConfigPath(configPath)
	c.setOpenAIKey(e.OpenAIKey)
	c.setGeminiKey(e.GeminiKey)
	c.setAllowEdit(e.AllowEdit)
	c.setAgent(e.Agent)
	c.setPromptType(prompt.DEFAULT)

	switch c.GetPromptType() {

	case prompt.DEFAULT:
		{
			c.setPrompt(prompt.NewDefaultPrompt())
		}
	case prompt.CUSTOM:
		{
			custom, err := prompt.NewCustomPrompt(prompt.CustomPromptDTO{
				Introduction: e.CustomPrompt.Introduction,
				Structure:    e.CustomPrompt.Structure,
				Rules:        e.CustomPrompt.Rules,
				Examples:     e.CustomPrompt.Examples,
				NewReader: func() (linereader.LineReader, error) {
					return readline.New("")
				},
				OutputWriter: os.Stdout,
				MenuAction:   prompt.NewMenuAction(),
				IsModify:     false,
			})

			if err != nil {
				return fmt.Errorf("Error on set Custom Prompt: %w", err)
			}

			c.setPrompt(custom)
		}
	default:
		{

			c.setPrompt(prompt.NewDefaultPrompt())
		}
	}

	return nil
}

func (c *Config) Config(reader io.Reader, agentOptions providers.AgentOptions, outputWriter io.Writer) error {
	fileConfig, err := os.OpenFile(c.configPath, os.O_RDWR, 0)

	if err != nil {
		return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
	}

	defer fileConfig.Close()

	f, err := io.ReadAll(fileConfig)

	if err != nil {
		return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
	}

	json.Unmarshal(f, &e)

	if err := c.configAgent(reader, agentOptions, outputWriter); err != nil {
		return err
	}

	c.configAllowEdit(reader, outputWriter)

	promptSelected, err := c.ConfigPrompt()

	if err != nil {
		return err
	}

	c.setPrompt(promptSelected)

	cfg := configDTO(fileConfig, c)

	if err = writeConfig(cfg); err != nil {
		return err
	}

	return nil
}

func (c *Config) ConfigPrompt() (prompt.Prompt, error) {
	menu := c.menuPromptSelector
	choice, err := menu.SelectPromptType()

	if err != nil {
		return nil, err
	}

	p := c.GetPrompt()

	switch choice {
	case prompt.CUSTOM:
		{
			c.setPromptType(prompt.CUSTOM)

			dto := prompt.CustomPromptDTO{
				Introduction: func() string {
					if p == nil {
						return ""
					}
					return p.GetIntroduction()
				}(),
				Structure: func() string {
					if p == nil {
						return ""
					}
					return p.GetStructure()
				}(),
				Rules: func() []string {
					if p == nil {
						return []string{}
					}
					return p.GetRules()
				}(),
				Examples: func() []string {
					if p == nil {
						return []string{}
					}
					return p.GetExamples()
				}(),
				NewReader: func() (linereader.LineReader, error) {
					return readline.New("")
				},
				OutputWriter: os.Stdout,
				MenuAction:   prompt.NewMenuAction(),
				IsModify:     true,
			}

			newPrompt, err := c.customPromptFactory.New(dto)

			if err != nil {
				return nil, err
			}

			return newPrompt, nil
		}

	case prompt.DEFAULT:
	default:
		{

			c.setPromptType(prompt.DEFAULT)
		}
	}

	return prompt.NewDefaultPrompt(), nil
}

func (c *Config) configAllowEdit(reader io.Reader, outputWriter io.Writer) {

	var userInputAllowEdit string
	options := []string{"y", "Y", "true", "n", "N", "false"}

	for !slices.Contains(options, userInputAllowEdit) {

		fmt.Fprintf(outputWriter, "Enable edit before make commit? %s ", "(y/n)")
		fmt.Fscanf(reader, "%s\n", &userInputAllowEdit)

		switch strings.ToLower(userInputAllowEdit) {
		case "y", "true":
			{
				c.setAllowEdit(true)
			}
		case "n", "false":
			{
				c.setAllowEdit(false)
			}
		default:
			{
				fmt.Fprintln(outputWriter, errs.InvalidEntryValue)
			}
		}
	}
}

func (c *Config) configAgent(reader io.Reader, agentOptions providers.AgentOptions, outputWriter io.Writer) error {

	var userInputOpenAi, userInputGemini string
	agentSelected := agentOptions.SelectedOption()

	switch agentSelected {
	case providers.OPEN_AI:
		{
			fmt.Fprint(outputWriter, "Write your OpenAI Key: ")
			fmt.Fscanf(reader, "%s\n", &userInputOpenAi)
			c.setOpenAIKey(userInputOpenAi)
			c.setAgent(providers.OPEN_AI)
		}
	case providers.GEMINI:
		{
			fmt.Fprint(outputWriter, "Write your Gemini Key: ")
			fmt.Fscanf(reader, "%s\n", &userInputGemini)
			c.setGeminiKey(userInputGemini)
			c.setAgent(providers.GEMINI)
		}
	default:
		{
			return errs.InvalidAgentSelected
		}
	}
	return nil
}

func convertJSON(data envStruct) ([]byte, error) {

	jsonByte, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return jsonByte, nil
}

type writeConfigDTO struct {
	file         *os.File
	gemini       string
	openai       string
	allowEdit    bool
	agent        providers.Provider
	promptType   prompt.PromptType
	customPrompt prompt.CustomPromptJSON
}

func configDTO(file *os.File, c *Config) *writeConfigDTO {
	p := c.GetPrompt()

	if p == nil {
		p = prompt.NewDefaultPrompt()
	}

	return &writeConfigDTO{
		file:       file,
		gemini:     c.GetGeminiKey(),
		openai:     c.GetOpenAIKey(),
		allowEdit:  c.GetAllowEdit(),
		agent:      c.GetAgent(),
		promptType: c.GetPromptType(),
		customPrompt: prompt.CustomPromptJSON{
			Introduction: p.GetIntroduction(),
			Structure:    p.GetStructure(),
			Examples:     p.GetExamples(),
			Rules:        p.GetRules(),
		},
	}
}

func writeConfig(v *writeConfigDTO) error {
	var customPrompt prompt.CustomPromptJSON

	if v.promptType == prompt.CUSTOM {
		customPrompt = v.customPrompt
	}

	data := envStruct{
		GeminiKey:    v.gemini,
		OpenAIKey:    v.openai,
		AllowEdit:    v.allowEdit,
		Agent:        v.agent,
		PromptType:   v.promptType,
		CustomPrompt: customPrompt,
	}

	dataByte, err := convertJSON(data)

	if err != nil {
		return fmt.Errorf(errs.ErrorOnConvertDataToByte+" -> %v", err)
	}

	if _, err := v.file.Seek(0, 0); err != nil {
		return fmt.Errorf("erro ao voltar o ponteiro: %v", err)
	}

	if err := v.file.Truncate(0); err != nil {
		return fmt.Errorf("erro ao truncar arquivo: %v", err)
	}

	if _, err = v.file.Write([]byte(dataByte)); err != nil {
		return fmt.Errorf(errs.ErrorOnWriteFileConfig+" -> %v", err)
	}

	return nil
}
