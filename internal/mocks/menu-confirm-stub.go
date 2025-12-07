package mocks

import "errors"

type StubMenuConfirm struct {
	Value bool
}

func (m *StubMenuConfirm) Run(label string) error {

	if m.Value {
		return nil
	}

	return errors.New("Canceled")

}

func NewStubMenuConfirm() *StubMenuConfirm {
	return &StubMenuConfirm{}
}
