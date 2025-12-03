package prompt

import (
	"fmt"

	"github.com/eduardongomes/gcai/internal/providers"
)

type mockMenu struct {
	OptionsToReturn []providers.MenuReturnOption
	Cursor          int
}

func (m *mockMenu) Display() (*providers.MenuReturnOption, error) {
	if m.Cursor >= len(m.OptionsToReturn) {
		return nil, fmt.Errorf("EOF: Finished menu options")
	}

	res := m.OptionsToReturn[m.Cursor]
	m.Cursor++

	return &res, nil
}

func (m *mockMenu) AddItem(l, v string) {}
