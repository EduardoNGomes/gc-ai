package providers

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"
)

func TestProviders(t *testing.T) {
	t.Run("Should testMenu", func(t *testing.T) {

		stub := mocks.NewStubMenuSelector()
		stub.Choices = append(stub.Choices, 1)
		agent := NewSelectAgent(stub)

		result, err := agent.SelectedOption()

		if err != nil {
		}

		if result != GEMINI {
			t.Fatalf("expected %s, got %s", GEMINI, result)
		}

		if len(stub.Items) != 2 {
			t.Fatalf("expected 2 menu items, got %d", len(stub.Items))
		}
	})

	t.Run("Shoud return selected option", func(t *testing.T) {
		agentSelected := NewSelectAgentSpy()

		agentSelected.SetAgent(OPEN_AI)

		option := agentSelected.SelectedOption()

		if option != OPEN_AI {
			t.Errorf("Expect %s, Receive %s", string(OPEN_AI), string(option))
		}
	})

	t.Run("Shoud return Gemini option", func(t *testing.T) {
		agentSelected := NewSelectAgentSpy()

		agentSelected.SetAgent(GEMINI)

		if agentSelected.option != GEMINI {
			t.Errorf("Expect %s, Receive %s", string(GEMINI), string(agentSelected.option))
		}
	})

	t.Run("Shoud return OPENAI option", func(t *testing.T) {
		agentSelected := NewSelectAgentSpy()

		agentSelected.SetAgent(OPEN_AI)

		if agentSelected.option != OPEN_AI {
			t.Errorf("Expect %s, Receive %s", string(OPEN_AI), string(agentSelected.option))
		}
	})

}
