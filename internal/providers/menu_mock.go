package providers

type MockMenu struct {
	items    []string
	toReturn string
}

func (m *MockMenu) AddItem(label, value string) {
	m.items = append(m.items, value)
}

func (m *MockMenu) Display() (*MenuReturnOption, error) {
	return &MenuReturnOption{Result: m.toReturn, Position: 0}, nil
}
