package providers

import "github.com/manifoldco/promptui"

type Menu interface {
	AddItem(label, value string)
	Display() string
}

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

func (m *ShellMenu) Display() string {
	prompt := promptui.Select{
		HideHelp:  true,
		Label:     m.title,
		Items:     m.items,
		IsVimMode: true,
		Size:      3,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return ""
	}

	return result
}
