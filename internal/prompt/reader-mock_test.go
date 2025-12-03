package prompt

import "errors"

type mockReader struct {
	Inputs []string
	Cursor int
}

func (m *mockReader) Readline() (string, error) {
	if m.Cursor >= len(m.Inputs) {
		return "", errors.New("EOF: Finished mock inputs")
	}
	res := m.Inputs[m.Cursor]
	m.Cursor++
	return res, nil
}

func (m *mockReader) Close() error                     { return nil }
func (m *mockReader) WriteStdin(b []byte) (int, error) { return len(b), nil }
func (m *mockReader) SetPrompt(s string)               {}
