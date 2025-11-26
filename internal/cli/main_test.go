package cli

import (
	"bytes"
	"testing"

	"github.com/eduardongomes/gcai/internal/agents"
	c "github.com/eduardongomes/gcai/internal/config"
	"github.com/eduardongomes/gcai/internal/flags"
	f "github.com/eduardongomes/gcai/internal/flags"
)

func TestCLI(t *testing.T) {

	flags := flags.Flags{
		OpenConfig:      false,
		EditCommit:      false,
		AlterEditConfig: nil,
	}

	t.Run("Should call config method on start method when config is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()
		agent := agents.NewMockAgent()

		cli := NewCLI(struct {
			Gemini agents.AgentMethods
			OpenAI agents.AgentMethods
		}{
			Gemini: agent,
			OpenAI: agent,
		})

		cli.Run(flags, confSpy, &bytes.Buffer{})

		expect := true
		result := confSpy.IsEmptyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})

	t.Run("Should call config key if confif is empty", func(t *testing.T) {
		confSpy := c.NewConfSpy()

		agent := agents.NewMockAgent()

		cli := NewCLI(struct {
			Gemini agents.AgentMethods
			OpenAI agents.AgentMethods
		}{
			Gemini: agent,
			OpenAI: agent,
		})

		cli.Run(flags, confSpy, &bytes.Buffer{})
		expect := true
		result := confSpy.IsConfigKeyCalled

		if result != expect {
			t.Errorf("Receive: '%v', Expect: '%v'", result, expect)
		}
	})

	t.Run("[Alter Config] should alter config be called", func(t *testing.T) {
		conf := c.NewConfSpy()
		agent := agents.NewMockAgent()

		cli := NewCLI(struct {
			Gemini agents.AgentMethods
			OpenAI agents.AgentMethods
		}{
			Gemini: agent,
			OpenAI: agent,
		})

		reader := &bytes.Buffer{}
		trueVal := true
		flags := f.Flags{
			AlterEditConfig: &trueVal,
		}

		cli.Run(flags, conf, reader)

		if conf.SetAllowEditCall == nil || *conf.SetAllowEditCall != true {
			t.Errorf("Expected SetAllowEdit to be called with true")
		}

		if agent.GetCommitCalled || agent.MakeCommitCalled {
			t.Errorf("Agents should not be called when altering config")
		}

	})

	t.Run("[Edit Commit] shoul call edit commit", func(t *testing.T) {
		agent := agents.NewMockAgent()
		cli := NewCLI(struct {
			Gemini agents.AgentMethods
			OpenAI agents.AgentMethods
		}{
			Gemini: agent,
			OpenAI: agent,
		})
		conf := c.NewConfSpy()
		reader := &bytes.Buffer{}

		flags := f.Flags{
			OpenConfig:      false,
			EditCommit:      true,
			AlterEditConfig: nil,
		}

		cli.Run(flags, conf, reader)

		if !agent.GetCommitCalled {
			t.Errorf("Expected GetCommit to be called")
		}
		if !agent.EditCalled {
			t.Errorf("Expected Edit to be called")
		}
		if !agent.MakeCommitCalled {
			t.Errorf("Expected MakeCommit to be called")
		}

	})

}
