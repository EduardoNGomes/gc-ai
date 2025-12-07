package config

import (
	"reflect"
	"testing"

	"github.com/eduardongomes/gcai/internal/prompt"
)

func TestConfigPrompt(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mockMenu := MockMenuSelector{
			Choice: prompt.CUSTOM,
		}

		expectedPrompt := &prompt.CustomPrompt{}

		mockFactory := &MockCustomPromptFactory{
			Returned: expectedPrompt,
		}

		c := &Config{
			menuPromptSelector:  mockMenu,
			customPromptFactory: mockFactory,
		}

		result, err := c.ConfigPrompt()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if result != expectedPrompt {
			t.Fatalf("expected %v, got %v", expectedPrompt, result)
		}

		if mockFactory.ReceivedDTO.IsModify != true {
			t.Fatal("Invalid dto passed to factory")
		}
	})

	t.Run("", func(t *testing.T) {
		mockMenu := MockMenuSelector{
			Choice: prompt.DEFAULT,
		}

		mockFactory := &MockCustomPromptFactory{}

		c := &Config{
			menuPromptSelector:  mockMenu,
			customPromptFactory: mockFactory,
		}

		result, err := c.ConfigPrompt()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expect := prompt.NewDefaultPrompt()

		if !reflect.DeepEqual(result, expect) {
			t.Fatalf("expected %v, got %v", expect, result)
		}
	})

}

type MockMenuSelector struct {
	Choice prompt.PromptType
	Err    error
}

func (m MockMenuSelector) SelectPromptType() (prompt.PromptType, error) {
	return m.Choice, m.Err
}

type MockCustomPromptFactory struct {
	ReceivedDTO prompt.CustomPromptDTO
	Returned    prompt.Prompt
	Err         error
}

func (m *MockCustomPromptFactory) New(dto prompt.CustomPromptDTO) (prompt.Prompt, error) {
	m.ReceivedDTO = dto
	return m.Returned, m.Err
}
