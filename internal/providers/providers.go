package providers

import "github.com/eduardongomes/gcai/internal/menu"

type Provider string

type SelectAgent struct {
	menu   menu.MenuSelector
	option Provider
}

const (
	OPEN_AI Provider = "OPEN_AI"
	GEMINI  Provider = "GEMINI"
)

type AgentOptions interface {
	SelectedOption() (Provider, error)
}

func (s *SelectAgent) SelectedOption() (Provider, error) {
	items := []string{
		string(OPEN_AI),
		string(GEMINI),
	}

	choice, err := s.menu.Run("Select your agent", items)

	if err != nil {
		return Provider(""), err
	}

	return Provider(choice.Result), nil
}

func NewSelectAgent() *SelectAgent {
	return &SelectAgent{menu: menu.NewProdMenuSelector()}
}
