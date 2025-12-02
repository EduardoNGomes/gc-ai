package linereader

type mockReader struct {
	input string
}

func (m *mockReader) WriteStdin(b []byte) (int, error) {
	m.input = string(b)
	return len(b), nil
}

func (m *mockReader) Readline() (string, error) {
	return m.input, nil
}

func (m *mockReader) Close() error {
	return nil
}

func (m *mockReader) SetPrompt(string) {}

func NewMockReader() *mockReader {
	return &mockReader{}
}
