package prompt

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

type Prompt interface {
	GetIntroduction() string
	GetRules() []string
	GetStructure() string
	GetExamples() []string
	ConvertToJSON() PromptJSON
	ConvertToPromptString() string
}
