package config

import "io"

type ConfigSpy struct {
	IsEmptyCalled     bool
	IsConfigKeyCalled bool
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
	return "test_key"
}

func (c *ConfigSpy) GetOpenAIKey() string {
	return "test_key"
}

func NewConfSpy() *ConfigSpy {
	return &ConfigSpy{}
}
