package mocks

import "errors"

type StubMenuWriter struct {
	Value string
}

func (m *StubMenuWriter) Run(label string) (string, error) {

	if m.Value == "error" {
		return "", errors.New("Mock ERROR")
	}

	return m.Value, nil

}

func NewStubMenuWriter() *StubMenuWriter {
	return &StubMenuWriter{}
}
