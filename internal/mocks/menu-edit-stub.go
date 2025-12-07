package mocks

import "errors"

type StubMenuEdit struct {
	Value string
}

func (m *StubMenuEdit) Run(label, defaultValue string) (string, error) {

	if m.Value == defaultValue {
		return "", errors.New("Canceled")
	}

	return m.Value, nil

}

func NewStubMenuEdit() *StubMenuEdit {
	return &StubMenuEdit{}
}
