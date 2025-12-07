package menu_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestMenuEdit(t *testing.T) {

	errValue := "error"
	value := "test"

	t.Run("[MenuEdit] Should return right value", func(t *testing.T) {
		menuEdit := mocks.NewStubMenuWriter()

		menuEdit.Value = value
		result, err := menuEdit.Run("test label")

		if err != nil {
			t.Errorf("Expect nil, receive %v", err)
		}

		if result != value {
			t.Errorf("Expect 'test 2', receive %v", result)
		}

	})

	t.Run("[MenuEdit] Should return err", func(t *testing.T) {
		menuEdit := mocks.NewStubMenuWriter()

		menuEdit.Value = errValue

		result, err := menuEdit.Run("test label")

		if err == nil {
			t.Errorf("Expect nil, receive %v", err)
		}

		if result != "" {
			t.Errorf("Expect empty string, receive %v", result)
		}

	})

}
