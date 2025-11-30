package providers

import "github.com/manifoldco/promptui"

type ShellMenu struct {
	title string
	items []string
}

func NewShellMenu(title string) *ShellMenu {
	return &ShellMenu{
		title: title,
		items: []string{},
	}
}

func (m *ShellMenu) AddItem(label, value string) {
	m.items = append(m.items, value)
}

func (m *ShellMenu) Display() (*MenuReturnOption, error) {
	prompt := promptui.Select{
		HideHelp: true,
		Label:    m.title,
		Items:    m.items,
		Size:     3,
	}

	i, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return &MenuReturnOption{Position: i, Result: result}, nil
}
