package prompt

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eduardongomes/gcai/internal/providers"
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

func NewMenuPromptOptions(m providers.Menu) (PromptType, error) {

	m.AddItem(string(DEFAULT), string(DEFAULT))
	m.AddItem(string(CUSTOM), string(CUSTOM))

	r, err := m.Display()

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
	rulesFormated := convertStringArrayPromptToString(p.GetRules())
	examplesFormated := convertStringArrayPromptToString(p.GetExamples())

	return fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", p.GetIntroduction(), rulesFormated, p.GetStructure(), examplesFormated)
}

func convertStringArrayPromptToString(arr []string) string {
	return "- " + strings.Join(arr, "\n- ")
}
