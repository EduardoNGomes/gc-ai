package providers_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/providers"
)

func TestProviders(t *testing.T) {

	t.Run("Shoud return selected option", func(t *testing.T) {
		agentSelected := providers.NewSelectAgentSpy()

		agentSelected.SetAgent(providers.OPEN_AI)

		option, err := agentSelected.SelectedOption()

		if err != nil {
			t.Errorf("Expected nil, receive %v", err)
		}

		if option != providers.OPEN_AI {
			t.Errorf("Expect %s, Receive %s", string(providers.OPEN_AI), string(option))
		}
	})

	t.Run("Shoud return Gemini option", func(t *testing.T) {
		agentSelected := providers.NewSelectAgentSpy()

		agentSelected.SetAgent(providers.GEMINI)

		if agentSelected.Option != providers.GEMINI {
			t.Errorf("Expect %s, Receive %s", string(providers.GEMINI), string(agentSelected.Option))
		}
	})

	t.Run("Shoud return OPENAI option", func(t *testing.T) {
		agentSelected := providers.NewSelectAgentSpy()

		agentSelected.SetAgent(providers.OPEN_AI)

		if agentSelected.Option != providers.OPEN_AI {
			t.Errorf("Expect %s, Receive %s", string(providers.OPEN_AI), string(agentSelected.Option))
		}
	})

}
