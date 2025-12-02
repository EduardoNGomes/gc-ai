package prompt

import (
	"testing"

	l "github.com/eduardongomes/gcai/internal/line-reader"
)

func TestCustomPrompt(t *testing.T) {

	customIndroduction := "Custom Introduction"
	customStructure := "Custom Structure"
	customExamples := []string{"Custom Examples"}
	customRules := []string{"Custom Rules"}
	reader := func() (l.LineReader, error) {
		return l.NewMockReader(), nil
	}

	t.Run("[GetIntroduction] Should return custom introduction", func(t *testing.T) {

		r, err := NewCustomPrompt(CustomPromptDTO{
			Introduction: customIndroduction,
			NewReader:    reader,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetIntroduction()
		expected := customIndroduction
		checkAssertString(t, result, expected)
	})
	t.Run("[GetStructure] Should return custom structure", func(t *testing.T) {
		sut, err := NewCustomPrompt(CustomPromptDTO{Structure: customStructure, NewReader: reader})

		if err != nil {
			t.Error(err)
		}
		result := sut.GetStructure()
		expected := customStructure

		checkAssertString(t, result, expected)
	})

	t.Run("[GetRules] Should return custom rules", func(t *testing.T) {

		r, err := NewCustomPrompt(CustomPromptDTO{
			Rules:     customRules,
			NewReader: reader,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetRules()
		expected := customRules

		checkAssertArray(t, result, expected)
	})
	t.Run("[GetExamples] Should return custom examples", func(t *testing.T) {

		r, err := NewCustomPrompt(CustomPromptDTO{
			Examples:  customExamples,
			NewReader: reader,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetExamples()

		expected := customExamples

		checkAssertArray(t, result, expected)
	})
}
