package providers

import "github.com/rrossmiller/gocliselect"

type Menu interface {
	AddItem(label, value string)
	Display() string
}

type ShellMenu struct {
	menu *gocliselect.Menu
}

func NewShellMenu(title string) *ShellMenu {
	shell := &ShellMenu{menu: gocliselect.NewMenu(title)}
	shell.menu.VimKeys = true
	return shell
}

func (m *ShellMenu) AddItem(label, value string) {
	m.menu.AddItem(label, value)
}

func (m *ShellMenu) Display() string {
	return m.menu.Display()
}
