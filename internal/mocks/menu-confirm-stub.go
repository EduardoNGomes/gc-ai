package mocks

import "errors"

type StubMenuConfirm struct {
	Values []bool
	cursor int
}

func (m *StubMenuConfirm) Run(label string) error {
	m.cursor = 0
	if m.Values[m.cursor] {
		m.cursor++
		return nil
	}

	return errors.New("Canceled")

}

func NewStubMenuConfirm() *StubMenuConfirm {
	return &StubMenuConfirm{}
}
