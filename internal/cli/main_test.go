package cli

import (
	"bytes"
	"testing"

	"github.com/eduardongomes/gcai/internal/agents"
	c "github.com/eduardongomes/gcai/internal/config"
	"github.com/eduardongomes/gcai/internal/flags"
)

func TestCLI(t *testing.T) {

	flags := flags.Flags{
		OpenConfig:      false,
		EditCommit:      false,
		AlterEditConfig: nil,
	}
	t.Run("Should call config method on start method when config is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		cli := NewCLI()
		agent := agents.NewMockAgent()

		cli.Run(flags, confSpy, agent, &bytes.Buffer{})

		expect := true
		result := confSpy.IsEmptyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})

	t.Run("Should call config key if confif is empty", func(t *testing.T) {
		cli := NewCLI()
		confSpy := c.NewConfSpy()

		agent := agents.NewMockAgent()
		cli.Run(flags, confSpy, agent, &bytes.Buffer{})
		expect := true
		result := confSpy.IsConfigKeyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})
}
