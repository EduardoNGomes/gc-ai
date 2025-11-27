package prompt

import (
	"testing"
)

func TestCustomPrompt(t *testing.T) {
	t.Run("[GetIntroduction] Should return custom introduction", func(t *testing.T) {
		customIndroduction := "Custom Introduction"
		result := NewCustomPrompt(customIndroduction, "", []string{}, []string{}).GetIntroduction()
		expected := customIndroduction
		checkAssertString(t, result, expected)
	})
	t.Run("[GetStructure] Should return custom structure", func(t *testing.T) {
		customStructure := "Custom Structure"
		result := NewCustomPrompt("", customStructure, []string{}, []string{}).GetStructure()
		expected := customStructure

		checkAssertString(t, result, expected)
	})
	t.Run("[GetRules] Should return custom rules", func(t *testing.T) {
		customRules := []string{"Custom Rules"}
		result := NewCustomPrompt("", "", customRules, []string{}).GetRules()
		expected := customRules

		checkAssertArray(t, result, expected)
	})
	t.Run("[GetExamples] Should return custom examples", func(t *testing.T) {
		customExamples := []string{"Custom Examples"}
		result := NewCustomPrompt("", "", []string{}, customExamples).GetExamples()
		expected := customExamples

		checkAssertArray(t, result, expected)
	})
}
