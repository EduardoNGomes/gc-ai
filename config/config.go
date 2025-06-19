package config

import (
	"github.com/gofor-little/env"
)

type Config struct {
	open_ai_key string
	gemini_key  string
}

func (c *Config) isEmpty() bool {
	gemini := c.getGeminiKey()
	openAI := c.getOpenIAKey()

	if len(gemini) == 0 && len(openAI) == 0 {
		return true
	}

	return false
}

func (c *Config) getOpenIAKey() string {
	return c.open_ai_key
}

func (c *Config) setOpenIAKey(v string) {
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

	c.setOpenIAKey(o)

	g, err := env.MustGet("GEMINI")

	if err != nil {
		env.Write("GEMINI", "", ".env", true)
	}

	c.setGeminiKey(g)

}
