package prompt

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

func NewCustomPrompt(introduction, structure string, rules, examples []string) *CustomPrompt {
	return &CustomPrompt{
		introduction,
		structure,
		rules,
		examples,
	}
}
