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

func (c *ConfigSpy) LoadEnvs() {}

func (c *ConfigSpy) ConfigKey(io.Reader) {
	c.IsConfigKeyCalled = true
}

func NewConfSpy() *ConfigSpy {
	return &ConfigSpy{}
}
