package mocks

import (
	"errors"

	"github.com/eduardongomes/gcai/internal/menu"
)

type StubMenuSelector struct {
	Choices []int
	cursor  int
	Items   []menu.MenuSelectorReturnOption
}

func (m *StubMenuSelector) Run(label string, items []string) (*menu.MenuSelectorReturnOption, error) {
	for i, v := range items {
		m.Items = append(m.Items, menu.MenuSelectorReturnOption{Position: i, Result: v})
	}

	if len(m.Items) < m.cursor+1 {
		return nil, errors.New("Invalid value entry")
	}

	position := m.Choices[m.cursor]

	if len(m.Items) < position+1 {
		return nil, errors.New("Invalid value entry")
	}

	result := &m.Items[position]

	m.cursor++
	return result, nil

}

func NewStubMenuSelector() *StubMenuSelector {
	return &StubMenuSelector{}
}
