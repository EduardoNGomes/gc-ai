package mocks

import "errors"

type StubMenuEditable struct {
	Value string
}

func (m *StubMenuEditable) Run(label, defaultValue string) (string, error) {

	if m.Value == defaultValue {
		return "", errors.New("Canceled")
	}

	return m.Value, nil

}

func NewStubMenuEditable() *StubMenuEditable {
	return &StubMenuEditable{}
}
