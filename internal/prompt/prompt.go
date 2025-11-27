package prompt

type Prompt interface {
	GetIntroduction() string
	GetRules() []string
	GetStructure() string
	GetExamples() []string
}
