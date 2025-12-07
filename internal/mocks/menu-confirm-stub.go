package mocks

import "errors"

type StubMenuConfirm struct {
	Values []bool
	cursor int
}

func (m *StubMenuConfirm) Run(label string) error {

	if m.cursor >= len(m.Values) {
		return errors.New("Not enough stub values provided")
	}

	result := m.Values[m.cursor]
	m.cursor++

	if result {
		return nil
	}

	return errors.New("Canceled")
}

func NewStubMenuConfirm() *StubMenuConfirm {
	return &StubMenuConfirm{}
}
