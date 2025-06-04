package config

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
func (c *Config) getGeminiKey() string {
	return c.gemini_key
}
