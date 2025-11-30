package providers

type Provider string

type SelectAgent struct {
	menu   Menu
	option Provider
}

const (
	OPEN_AI Provider = "OPEN_AI"
	GEMINI  Provider = "GEMINI"
)

type AgentOptions interface {
	SelectedOption() Provider
}

func (s *SelectAgent) SelectedOption() Provider {
	s.menu.AddItem(string(OPEN_AI), string(OPEN_AI))
	s.menu.AddItem(string(GEMINI), string(GEMINI))

	choice, _ := s.menu.Display()
	return Provider(choice.Result)
}

func NewSelectAgent(m Menu) *SelectAgent {
	return &SelectAgent{menu: m}
}
