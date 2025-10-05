package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/eduardongomes/gcai/errs"
)

type Config struct {
	open_ai_key string
	gemini_key  string
	configPath  string
	allow_edit  bool
}

type envStruct struct {
	GeminiKey string `json:"geminiKey"`
	OpenAIKey string `json:"openAIAKey"`
	AllowEdit bool   `json:"allowEdit"`
}

var e envStruct

type ConfigMethods interface {
	LoadEnvs(configPath string) error
	IsEmpty() bool
	ConfigKey(io.Reader) error
	GetGeminiKey() string
	GetOpenAIKey() string
	GetAllowEdit() bool
	SetAllowEdit(v, rewrite bool) error
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

	if rewrite == true {
		fileConfig, err := os.OpenFile(c.configPath, os.O_RDWR, 0)

		if err != nil {
			return fmt.Errorf(errs.CannotOpenFileErr+" -> %w", err)
		}

		defer fileConfig.Close()

		if err = writeConfig(fileConfig, c.gemini_key, c.open_ai_key, c.allow_edit); err != nil {
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

		writeConfig(fileConfig, "", "", false)
	}

	c.setConfigPath(configPath)
	c.setOpenAIKey(e.OpenAIKey)
	c.setGeminiKey(e.GeminiKey)
	c.SetAllowEdit(e.AllowEdit, false)
	return nil
}

func (c *Config) ConfigKey(reader io.Reader) error {
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

	var useInputOpenAi, useInputGemini string

	fmt.Print("Write your OpenAI Key: ")
	fmt.Fscanf(reader, "%s\n", &useInputOpenAi)
	c.setOpenAIKey(useInputOpenAi)

	fmt.Print("Write your Gemini Key: ")
	fmt.Fscanf(reader, "%s\n", &useInputGemini)
	c.setGeminiKey(useInputGemini)

	if err = writeConfig(fileConfig, c.gemini_key, c.open_ai_key, c.allow_edit); err != nil {
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

func writeConfig(f *os.File, geminiV, openAIV string, allowEdit bool) error {
	data := envStruct{
		GeminiKey: geminiV,
		OpenAIKey: openAIV,
		AllowEdit: allowEdit,
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
