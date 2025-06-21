package config

import "io"

type ConfigSpy struct {
	isEmptyCalled bool
}

func (c *ConfigSpy) isEmpty() bool {
	c.isEmptyCalled = true

	return true
}

func (c *ConfigSpy) getGeminiKey() string {
	return ""
}

func (c *ConfigSpy) setGeminiKey(string) {}

func (c *ConfigSpy) getOpenAIKey() string {
	return ""
}

func (c *ConfigSpy) setOpenAIKey(string) {}

func (c *ConfigSpy) loadEnvs() {}

func (c *ConfigSpy) configKey(io.Reader) {}

func newConfSpy() *ConfigSpy {
	return &ConfigSpy{}
}
