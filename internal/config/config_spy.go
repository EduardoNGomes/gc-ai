package config

import (
	"github.com/eduardongomes/gcai/internal/prompt"
	"github.com/eduardongomes/gcai/internal/providers"
)

type ConfigSpy struct {
	IsEmptyCalled       bool
	IsConfigKeyCalled   bool
	geminiKey           string
	openAIKey           string
	allow_edit          bool
	agent               providers.Provider
	SetAllowEditCall    *bool
	prompt              prompt.Prompt
	promptType          prompt.PromptType
	RewriteConfigCalled bool
	ConfigCalled        bool
	IsEmptyResult       bool
	GetAgentReturn      providers.Provider
	GetAllowEditReturn  bool
}

func (c *ConfigSpy) IsEmpty() bool {
	c.IsEmptyCalled = true
	return true
}

func (c *ConfigSpy) LoadEnvs(path string) error {
	return nil
}

func (c *ConfigSpy) Config(providers.AgentOptions) error {
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

func (c *ConfigSpy) SetOpenAIKey(i string) {
	c.openAIKey = i
}

func (c *ConfigSpy) GetAllowEdit() bool {
	return c.allow_edit
}

func (c *ConfigSpy) setAllowEdit(val bool) {
	c.SetAllowEditCall = &val
}

func (c *ConfigSpy) GetAgent() providers.Provider {
	return c.agent
}

func (c *ConfigSpy) setAgent(v providers.Provider) {
	c.agent = v
}

func (c *ConfigSpy) GetPromptString() string {
	return prompt.ConvertToPromptString(c.prompt)
}

func (c *ConfigSpy) GetPrompt() prompt.Prompt {
	return c.prompt
}

func (c *ConfigSpy) setPrompt(v prompt.Prompt) {
	c.prompt = v
}

func (c *ConfigSpy) GetPromptType() prompt.PromptType {
	return c.promptType
}

func (c *ConfigSpy) setPromptType(v prompt.PromptType) {
	c.promptType = v
}

func (c *ConfigSpy) RewriteConfig(v RewriteConfigOptions) error {
	if v.Agent != nil {
		c.setAgent(*v.Agent)
	}
	if v.AllowEdit != nil {
		c.setAllowEdit(*v.AllowEdit)
	}
	if v.PromptType != nil {
		c.setPromptType(*v.PromptType)
	}
	if v.Prompt != nil {
		c.setPrompt(*v.Prompt)
	}
	c.RewriteConfigCalled = true
	return nil
}

func NewConfSpy() *ConfigSpy {
	return &ConfigSpy{
		geminiKey:  "test_key",
		openAIKey:  "test_key",
		agent:      providers.GEMINI,
		allow_edit: false,
		promptType: prompt.DEFAULT,
		prompt:     prompt.NewDefaultPrompt(),
	}
}
