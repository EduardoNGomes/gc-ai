package prompt

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CustomPrompt struct {
	introduction string
	structure    string
	rules        []string
	examples     []string
}

func (p *CustomPrompt) GetIntroduction() string {
	return p.introduction
}

func (p *CustomPrompt) GetStructure() string {
	return p.structure
}

func (p *CustomPrompt) GetRules() []string {
	return p.rules
}

func (p *CustomPrompt) GetExamples() []string {
	return p.examples
}

func (p *CustomPrompt) ConvertToJSON() (string, error) {
	data := PromptJSON{
		Introduction: p.GetIntroduction(),
		Rules:        p.GetRules(),
		Structure:    p.GetStructure(),
		Examples:     p.GetExamples(),
	}

	json, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(json), nil
}

func (p *CustomPrompt) ConvertToPromptString() string {
	rulesFormated := "- " + strings.Join(p.GetRules(), "\n- ")
	examplesFormated := "- " + strings.Join(p.GetExamples(), "\n- ")

	return fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", p.GetIntroduction(), rulesFormated, p.GetStructure(), examplesFormated)
}

func NewCustomPrompt(introduction, structure string, rules, examples []string) *CustomPrompt {
	return &CustomPrompt{
		introduction,
		structure,
		rules,
		examples,
	}
}
