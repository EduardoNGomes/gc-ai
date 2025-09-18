//go:build integration
// +build integration

package agents

import (
	"testing"

	c "github.com/eduardongomes/gcai/internal/config"
)

func TestAgent(t *testing.T) {
	t.Run("Should receive error if has wrong api_key", func(t *testing.T) {
		agent := NewGeminiAgent()
		c := c.NewConfSpy()

		diff := `diff --git a/main.go b/main.go
   				 +fmt.Println("hello world")`

		_, err := agent.GetCommit(c, diff)

		if err == nil {
			t.Errorf("Should not work with fake key")
		}

	})

}
