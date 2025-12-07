package menu_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestMenuSelector(t *testing.T) {
	t.Run("[MenuSelector] Should return right value", func(t *testing.T) {
		menuSelector := mocks.NewStubMenuSelector()
		items := []string{
			"test 1",
			"test 2",
			"test 3",
			"test 4",
		}
		menuSelector.Choices = append(menuSelector.Choices, 1)

		result, err := menuSelector.Run("test label", items)

		if err != nil {
			t.Errorf("Expect nil, receive %v", err)
		}
		if result.Result != "test 2" {

			t.Errorf("Expect 'test 2', receive %v", result)
		}

	})

	t.Run("[MenuSelector] Should return err", func(t *testing.T) {
		menuSelector := mocks.NewStubMenuSelector()
		items := []string{
			"test 1",
		}
		menuSelector.Choices = append(menuSelector.Choices, 1)

		result, err := menuSelector.Run("test label", items)

		if err == nil {
			t.Errorf("Expect nil, receive %v", result)
		}

	})

}
