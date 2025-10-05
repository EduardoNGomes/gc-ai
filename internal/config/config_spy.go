package config

import "io"

type ConfigSpy struct {
	IsEmptyCalled     bool
	IsConfigKeyCalled bool
	geminiKey         string
	openAIKey         string
	allow_edit        bool
}

func (c *ConfigSpy) IsEmpty() bool {
	c.IsEmptyCalled = true
	return true
}

func (c *ConfigSpy) LoadEnvs(path string) error {
	return nil
}

func (c *ConfigSpy) ConfigKey(io.Reader) error {
	c.IsConfigKeyCalled = true
	return nil
}

func (c *ConfigSpy) GetGeminiKey() string {
	return c.geminiKey
}

func (c *ConfigSpy) SetGeminiKey(i string) {
	c.geminiKey = i
}

func (c *ConfigSpy) GetOpenAIKey() string {
	return c.openAIKey
}

func (c *ConfigSpy) SetOpenAiKey(i string) {
	c.openAIKey = i
}

func (c *ConfigSpy) GetAllowEdit() bool {
	return c.allow_edit
}

func (c *ConfigSpy) SetAllowEdit(v, rewrite bool) error {
	return nil
}

func NewConfSpy() *ConfigSpy {
	return &ConfigSpy{
		geminiKey: "test_key",
		openAIKey: "test_key",
	}
}
