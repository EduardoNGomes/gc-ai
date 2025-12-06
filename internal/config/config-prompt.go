package config

import (
	"github.com/eduardongomes/gcai/internal/prompt"
	"github.com/eduardongomes/gcai/internal/providers"
)

type MenuPromptSelector interface {
	SelectPromptType() (prompt.PromptType, error)
}

type ProdMenuSelector struct{}

func (ProdMenuSelector) SelectPromptType() (prompt.PromptType, error) {
	menu := providers.NewMenuCustomPrompt("Choice your prompt type")
	return prompt.NewMenuPromptOptions(menu)
}

type CustomPromptFactory interface {
	New(dto prompt.CustomPromptDTO) (prompt.Prompt, error)
}

type ProdCustomPromptFactory struct{}

func (ProdCustomPromptFactory) New(dto prompt.CustomPromptDTO) (prompt.Prompt, error) {
	return prompt.NewCustomPrompt(dto)
}
