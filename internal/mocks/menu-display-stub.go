package mocks

import (
	"errors"

	"github.com/eduardongomes/gcai/internal/menu"
)

type StubMenuDisplay struct {
	Cursor int
	items  []menu.MenuSelectorReturnOption
}

func (m *StubMenuDisplay) Run(label string, items []string) (*menu.MenuSelectorReturnOption, error) {
	for i, v := range items {
		m.items = append(m.items, menu.MenuSelectorReturnOption{Position: i, Result: v})
	}

	if len(m.items) < m.Cursor+1 {
		return nil, errors.New("Invalid value entry")
	}

	return &m.items[m.Cursor], nil

}

func NewStubMenuDisplay() *StubMenuDisplay {
	return &StubMenuDisplay{}
}
