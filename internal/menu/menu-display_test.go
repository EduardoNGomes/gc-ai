package menu_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestMenuDisplay(t *testing.T) {
	t.Run("[MenuDisplay] Should return right value", func(t *testing.T) {
		menuDisplay := mocks.NewStubMenuDisplay()
		items := []string{
			"test 1",
			"test 2",
			"test 3",
			"test 4",
		}
		menuDisplay.Cursor = 1

		result, err := menuDisplay.Run("test label", items)

		if err != nil {
			t.Errorf("Expect nil, receive %v", err)
		}
		if result.Result != "test 2" {

			t.Errorf("Expect 'test 2', receive %v", result)
		}

	})

	t.Run("[MenuDisplay] Should return err", func(t *testing.T) {
		menuDisplay := mocks.NewStubMenuDisplay()
		items := []string{
			"test 1",
		}
		menuDisplay.Cursor = 1

		result, err := menuDisplay.Run("test label", items)

		if err == nil {
			t.Errorf("Expect nil, receive %v", result)
		}

	})

}
