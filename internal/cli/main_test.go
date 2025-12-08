package cli

import (
	"bytes"
	"testing"

	"github.com/eduardongomes/gcai/internal/agents"
	c "github.com/eduardongomes/gcai/internal/config"
	f "github.com/eduardongomes/gcai/internal/flags"
	"github.com/eduardongomes/gcai/internal/menu"
	"github.com/eduardongomes/gcai/internal/providers"
)

func TestCLI(t *testing.T) {

	defaultFlags := f.Flags{
		OpenConfig:      false,
		EditCommit:      false,
		AlterEditConfig: nil,
		AlterAgent:      nil,
	}

	setupSafeCLI := func(agent agents.AgentMethods) *CLI {
		cli := NewCLI(struct {
			Gemini agents.AgentMethods
			OpenAI agents.AgentMethods
		}{
			Gemini: agent,
			OpenAI: agent,
		})

		cli.Committer = func(msg string) error {
			return nil
		}

		cli.Editter = func(m menu.MenuEditable, msg string) (string, error) {
			return msg, nil
		}

		return cli
	}

	t.Run("Should call config method on start method when config is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		confSpy.IsEmptyResult = true

		agent := agents.NewMockAgent()
		agent.GetCommitReturn = "test commit"

		cli := setupSafeCLI(agent)

		cli.Run(defaultFlags, confSpy, &bytes.Buffer{})

		if !confSpy.IsEmptyCalled {
			t.Errorf("Expected IsEmpty() to be called")
		}

	})

	t.Run("Should call config key if config is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		confSpy.IsEmptyResult = true

		agent := agents.NewMockAgent()
		agent.GetCommitReturn = "test commit"

		cli := setupSafeCLI(agent)

		cli.Run(defaultFlags, confSpy, &bytes.Buffer{})

		if !confSpy.IsConfigKeyCalled {
			t.Errorf("Expected Config setup to be called")
		}
	})

	t.Run("[Alter Config] should alter config be called", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		agent := agents.NewMockAgent()

		cli := setupSafeCLI(agent)

		trueVal := true
		flags := f.Flags{
			AlterEditConfig: &trueVal,
		}

		cli.Run(flags, confSpy, &bytes.Buffer{})

		if !confSpy.RewriteConfigCalled {
			t.Errorf("Expected RewriteConfig to be called")
		}

		if agent.GetCommitCalled {
			t.Errorf("Agents should not be called when altering config")
		}
	})

	t.Run("[Edit Commit] should call edit commit", func(t *testing.T) {
		agent := agents.NewMockAgent()
		agent.GetCommitReturn = "Initial Msg"

		cli := setupSafeCLI(agent)

		editCalled := false
		commitCalled := false
		finalMsg := ""

		cli.Editter = func(m menu.MenuEditable, msg string) (string, error) {
			editCalled = true
			return "Edited Msg", nil
		}

		cli.Committer = func(msg string) error {
			commitCalled = true
			finalMsg = msg
			return nil
		}

		confSpy := c.NewConfSpy()
		confSpy.GetAgentReturn = providers.GEMINI
		confSpy.GetAllowEditReturn = false

		flags := f.Flags{
			EditCommit: true,
		}

		cli.Run(flags, confSpy, &bytes.Buffer{})

		if !editCalled {
			t.Error("Expected Editter to be called")
		}
		if !commitCalled {
			t.Error("Expected Committer to be called")
		}
		if finalMsg != "Edited Msg" {
			t.Errorf("Expected 'Edited Msg', got '%s'", finalMsg)
		}
	})
}
