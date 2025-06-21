package config

import (
	"fmt"
	"io"

	"github.com/gofor-little/env"
)

type Config struct {
	open_ai_key string
	gemini_key  string
}

type ConfigMethods interface {
	isEmpty() bool

	getGeminiKey() string
	setGeminiKey(string)

	getOpenAIKey() string
	setOpenAIKey(string)

	loadEnvs()

	configKey(io.Reader)
}

func (c *Config) Start(config ConfigMethods) {
	c.loadEnvs()

}

func (c *Config) isEmpty() bool {
	gemini := c.getGeminiKey()
	openAI := c.getOpenAIKey()

	if len(gemini) == 0 && len(openAI) == 0 {
		return true
	}

	return false
}

func (c *Config) getOpenAIKey() string {
	return c.open_ai_key
}

func (c *Config) setOpenAIKey(v string) {
	c.open_ai_key = v
}

func (c *Config) getGeminiKey() string {
	return c.gemini_key
}

func (c *Config) setGeminiKey(v string) {
	c.gemini_key = v
}

func (c *Config) loadEnvs() {
	if err := env.Load("../.env"); err != nil {
		panic(err)
	}

	o, err := env.MustGet("OPEN_AI")

	if err != nil {
		env.Write("OPEN_AI", "", ".env", true)
	}

	c.setOpenAIKey(o)

	g, err := env.MustGet("GEMINI")

	if err != nil {
		env.Write("GEMINI", "", ".env", true)
	}

	c.setGeminiKey(g)
}

func (c *Config) configKey(reader io.Reader) {
	var useInputOpenAi, useInputGemini string

	fmt.Print("Write your OpenAI Key: ")
	fmt.Fscanf(reader, "%s\n", &useInputOpenAi)
	c.setOpenAIKey(useInputOpenAi)

	fmt.Print("Write your Gemini Key: ")
	fmt.Fscanf(reader, "%s\n", &useInputGemini)
	c.setGeminiKey(useInputGemini)
}

func NewConfig() *Config {
	return &Config{}
}
