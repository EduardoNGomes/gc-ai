package mocks

import (
	"errors"

	"github.com/eduardongomes/gcai/internal/menu"
)

type StubMenuSelector struct {
	Cursor int
	Items  []menu.MenuSelectorReturnOption
}

func (m *StubMenuSelector) Run(label string, items []string) (*menu.MenuSelectorReturnOption, error) {
	for i, v := range items {
		m.Items = append(m.Items, menu.MenuSelectorReturnOption{Position: i, Result: v})
	}

	if len(m.Items) < m.Cursor+1 {
		return nil, errors.New("Invalid value entry")
	}

	return &m.Items[m.Cursor], nil

}

func NewStubMenuSelector() *StubMenuSelector {
	return &StubMenuSelector{}
}
