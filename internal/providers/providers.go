package providers

import "github.com/rrossmiller/gocliselect"

type Provider string

type SelectAgent struct {
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
	menu := gocliselect.NewMenu("Chose a Agent")
	menu.VimKeys = true

	menu.AddItem(string(OPEN_AI), string(OPEN_AI))
	menu.AddItem(string(GEMINI), string(GEMINI))

	choice := menu.Display()

	return Provider(choice)
}

func NewSelectAgent() *SelectAgent {
	return &SelectAgent{}
}
