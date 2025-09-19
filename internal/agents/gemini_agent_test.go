package agents

import (
	"errors"
	"testing"

	"github.com/eduardongomes/gcai/errs"
	c "github.com/eduardongomes/gcai/internal/config"
)

func TestGeminiAgent(t *testing.T) {
	t.Run("Should receive error if has wrong api_key", func(t *testing.T) {
		agent := NewGeminiAgent()
		c := c.NewConfSpy()

		c.SetGeminiKey("")

		_, err := agent.GetCommit(c)

		if err == nil {
			t.Errorf("Should not work with fake key")
		}

		if !errors.Is(err, errs.EmptyKeyError) {
			t.Errorf("Should not work without empty key %v", err)
		}

	})

}
