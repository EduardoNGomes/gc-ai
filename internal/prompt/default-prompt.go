package prompt

import (
	"encoding/json"
	"fmt"
	"strings"
)

type DefaultPrompt struct {
	introduction string
	structure    string
	rules        []string
	examples     []string
}

const INTRODUCTION = "You are an expert software engineer trained in the Conventional Commits standard.Analyze the following git diff and produce a single-line commit message using this exact rules,struct and exemple:"

const STRUCTURE = "type: message"

var EXAMPLES = []string{
	"feat: implement login handler for member authentication",
	"feat: add internal link suggestion",
	"fix: handle 404 errors for user profiles",
	"refactor: abstract parsing logic into `Parser` class",
}

var RULES = []string{
	"Only output the final commit message",
	"Type MUST be lowercase",
	"Scope is NOT allowed (do NOT use parentheses)",
	"Message must be concise and descriptive",
	"Prefer referencing identifiers using code spans (e.g., `AuthService`)",
	"No emojis, no extra commentary, no multi-line output",
	"Avoid vague statements like 'update code' or 'misc changes'",
}

func (p *DefaultPrompt) GetIntroduction() string {
	return p.introduction
}

func (p *DefaultPrompt) GetStructure() string {
	return p.structure
}

func (p *DefaultPrompt) GetRules() []string {
	return p.rules
}

func (p *DefaultPrompt) GetExamples() []string {
	return p.examples
}

func (p *DefaultPrompt) ConvertToJSON() (string, error) {
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

func (p *DefaultPrompt) ConvertToPromptString() string {
	rulesFormated := "- " + strings.Join(p.GetRules(), "\n- ")
	examplesFormated := "- " + strings.Join(p.GetExamples(), "\n- ")

	return fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", p.GetIntroduction(), rulesFormated, p.GetStructure(), examplesFormated)
}

func NewDefaultPrompt() *DefaultPrompt {
	return &DefaultPrompt{
		introduction: INTRODUCTION,
		structure:    STRUCTURE,
		rules:        RULES,
		examples:     EXAMPLES,
	}
}
