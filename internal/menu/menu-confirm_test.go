package menu_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestMenuConfirm(t *testing.T) {
	t.Run("[MenuConfirm] Should return nil", func(t *testing.T) {
		menuConfirm := mocks.NewStubMenuConfirm()
		menuConfirm.Value = true

		if err := menuConfirm.Run("test label"); err != nil {
			t.Errorf("Expect nil, receive %v", err)
		}

	})

	t.Run("[MenuConfirm] Should return err", func(t *testing.T) {
		menuConfirm := mocks.NewStubMenuConfirm()
		menuConfirm.Value = false

		if err := menuConfirm.Run("test label"); err == nil {
			t.Errorf("Expect nil, receive %v", err)
		}

	})

}
