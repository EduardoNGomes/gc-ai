package prompt

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eduardongomes/gcai/internal/menu"
)

type PromptType string

const (
	DEFAULT PromptType = "DEFAULT"
	CUSTOM  PromptType = "CUSTOM"
)

type PromptJSON struct {
	Introduction string   `json:"introduction"`
	Rules        []string `json:"rules"`
	Structure    string   `json:"structure"`
	Examples     []string `json:"examples"`
}

func NewMenuPromptOptions(m menu.MenuSelector) (PromptType, error) {

	items := []string{
		string(DEFAULT),
		string(CUSTOM),
	}

	r, err := m.Run("Choice your prompt type", items)

	if err != nil {
		return "", err
	}

	return PromptType(r.Result), nil
}

type Prompt interface {
	GetIntroduction() string
	GetRules() []string
	GetStructure() string
	GetExamples() []string
}

func ConvertToJSON(p Prompt) ([]byte, error) {
	data := PromptJSON{
		Introduction: p.GetIntroduction(),
		Rules:        p.GetRules(),
		Structure:    p.GetStructure(),
		Examples:     p.GetExamples(),
	}

	json, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return json, nil
}

func ConvertToPromptString(p Prompt) string {
	rulesFormated := ConvertStringArrayPromptToString(p.GetRules())
	examplesFormated := ConvertStringArrayPromptToString(p.GetExamples())

	return fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", p.GetIntroduction(), rulesFormated, p.GetStructure(), examplesFormated)
}

func ConvertStringArrayPromptToString(arr []string) string {
	return "- " + strings.Join(arr, "\n- ")
}
