package providers

import "github.com/manifoldco/promptui"

type MenuCustomPrompt struct {
	title string
	items []string
}

func NewMenuCustomPrompt(title string) *MenuCustomPrompt {
	return &MenuCustomPrompt{
		title: title,
		items: []string{},
	}
}

func (m *MenuCustomPrompt) AddItem(label, value string) {
	m.items = append(m.items, value)
}

func (m *MenuCustomPrompt) Display() (*MenuReturnOption, error) {
	prompt := promptui.Select{
		HideHelp: true,
		Label:    m.title,
		Items:    m.items,
		Size:     6,
	}

	i, result, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	return &MenuReturnOption{Position: i, Result: result}, nil
}
