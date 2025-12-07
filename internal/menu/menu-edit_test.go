package menu_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestMenuEdit(t *testing.T) {
	defaultValue := "default value test"
	newValue := "new value test"

	t.Run("[MenuEdit] Should return right value", func(t *testing.T) {
		menuEdit := mocks.NewStubMenuEdit()

		menuEdit.Value = newValue
		result, err := menuEdit.Run("test label", defaultValue)

		if err != nil {
			t.Errorf("Expect nil, receive %v", err)
		}

		if result != newValue {
			t.Errorf("Expect 'test 2', receive %v", result)
		}

	})

	t.Run("[MenuEdit] Should return err", func(t *testing.T) {
		menuEdit := mocks.NewStubMenuEdit()

		menuEdit.Value = defaultValue

		result, err := menuEdit.Run("test label", defaultValue)

		if err == nil {
			t.Errorf("Expect nil, receive %v", err)
		}

		if result != "" {
			t.Errorf("Expect empty string, receive %v", result)
		}

	})

}
