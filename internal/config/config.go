package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/prompt"
	"github.com/eduardongomes/gcai/internal/providers"
)

type Config struct {
	open_ai_key string
	gemini_key  string
	configPath  string
	allow_edit  bool
	agent       providers.Provider
	promptType  prompt.PromptType
	prompt      prompt.Prompt
}

type envStruct struct {
	GeminiKey  string             `json:"geminiKey"`
	OpenAIKey  string             `json:"openAIAKey"`
	AllowEdit  bool               `json:"allowEdit"`
	Agent      providers.Provider `json:"agent"`
	PromptType prompt.PromptType  `json:"promptType"`
	Prompt     string             `json:"prompt"`
}

var e envStruct

type ConfigMethods interface {
	LoadEnvs(configPath string) error
	IsEmpty() bool
	ConfigKey(io.Reader, providers.AgentOptions, io.Writer) error

	GetGeminiKey() string
	GetOpenAIKey() string

	GetAllowEdit() bool
	SetAllowEdit(v, rewrite bool) error

	GetAgent() providers.Provider
	SetAgent(providers.Provider, bool) error

	GetPromptType() prompt.PromptType
	setPromptType(prompt.PromptType)

	GetPrompt() string
	setPrompt(prompt.Prompt)
}

func NewConfig() *Config {
	return &Config{}
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

func (c *Config) SetAllowEdit(v, rewrite bool) error {
	c.allow_edit = v

	if rewrite {
		fileConfig, err := os.OpenFile(c.configPath, os.O_RDWR, 0)

		if err != nil {
			return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
		}

		defer fileConfig.Close()

		if err = writeConfig(fileConfig, c.GetGeminiKey(), c.GetOpenAIKey(), c.GetAllowEdit(), c.GetAgent()); err != nil {
			return err
		}
	}

	return nil
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

func (c *Config) SetAgent(v providers.Provider, rewrite bool) error {
	c.agent = v

	if rewrite {
		fileConfig, err := os.OpenFile(c.configPath, os.O_RDWR, 0)

		if err != nil {
			return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
		}

		defer fileConfig.Close()

		if err = writeConfig(fileConfig, c.GetGeminiKey(), c.GetOpenAIKey(), c.GetAllowEdit(), c.GetAgent()); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) GetPrompt() string {
	return prompt.ConvertToPromptString(c.prompt)
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

		writeConfig(fileConfig, "", "", false, providers.GEMINI)
	}

	c.setConfigPath(configPath)
	c.setOpenAIKey(e.OpenAIKey)
	c.setGeminiKey(e.GeminiKey)
	c.SetAllowEdit(e.AllowEdit, false)
	c.SetAgent(e.Agent, false)
	return nil
}

func (c *Config) ConfigKey(reader io.Reader, agentOptions providers.AgentOptions, outputWriter io.Writer) error {
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

	var userInputOpenAi, userInputGemini, userInputAllowEdit string

	agentSelected := agentOptions.SelectedOption()

	switch agentSelected {
	case providers.OPEN_AI:
		{
			fmt.Fprint(outputWriter, "Write your OpenAI Key: ")
			fmt.Fscanf(reader, "%s\n", &userInputOpenAi)
			c.setOpenAIKey(userInputOpenAi)
			c.SetAgent(providers.OPEN_AI, false)
		}
	case providers.GEMINI:
		{
			fmt.Fprint(outputWriter, "Write your Gemini Key: ")
			fmt.Fscanf(reader, "%s\n", &userInputGemini)
			c.setGeminiKey(userInputGemini)
			c.SetAgent(providers.GEMINI, false)
		}
	default:
		{
			return errs.InvalidAgentSelected
		}
	}

	options := []string{"y", "Y", "true", "n", "N", "false"}

	for !slices.Contains(options, userInputAllowEdit) {

		fmt.Fprintf(outputWriter, "Enable edit before make commit? %s ", "(y/n)")
		fmt.Fscanf(reader, "%s\n", &userInputAllowEdit)

		switch strings.ToLower(userInputAllowEdit) {
		case "y", "true":
			{
				c.SetAllowEdit(true, false)
			}
		case "n", "false":
			{
				c.SetAllowEdit(false, false)
			}
		default:
			{
				fmt.Fprintln(outputWriter, errs.InvalidEntryValue)
			}
		}
	}

	if err = writeConfig(fileConfig, c.GetGeminiKey(), c.GetOpenAIKey(), c.GetAllowEdit(), c.GetAgent()); err != nil {
		return err
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

func writeConfig(f *os.File, geminiV, openAIV string, allowEdit bool, agent providers.Provider) error {
	data := envStruct{
		GeminiKey: geminiV,
		OpenAIKey: openAIV,
		AllowEdit: allowEdit,
		Agent:     agent,
	}

	dataByte, err := convertJSON(data)

	if err != nil {
		return fmt.Errorf(errs.ErrorOnConvertDataToByte+" -> %v", err)
	}

	if _, err := f.Seek(0, 0); err != nil {
		return fmt.Errorf("erro ao voltar o ponteiro: %v", err)
	}

	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("erro ao truncar arquivo: %v", err)
	}

	if _, err = f.Write([]byte(dataByte)); err != nil {
		return fmt.Errorf(errs.ErrorOnWriteFileConfig+" -> %v", err)
	}

	return nil
}
