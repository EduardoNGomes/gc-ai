package prompt

import (
	"reflect"
	"testing"
)

func TestDefaultPrompt(t *testing.T) {
	t.Run("[GetIntroduction] Should return default introduction", func(t *testing.T) {
		result := NewDefaultPrompt().GetIntroduction()
		expected := INTRODUCTION
		checkAssertString(t, result, expected)
	})
	t.Run("[GetStructure] Should return default structure", func(t *testing.T) {
		result := NewDefaultPrompt().GetStructure()
		expected := STRUCTURE
		checkAssertString(t, result, expected)
	})
	t.Run("[GetRules] Should return default rules", func(t *testing.T) {
		result := NewDefaultPrompt().GetRules()
		expected := RULES
		checkAssertArray(t, result, expected)
	})
	t.Run("[GetExamples] Should return default examples", func(t *testing.T) {
		result := NewDefaultPrompt().GetExamples()
		expected := EXAMPLES
		checkAssertArray(t, result, expected)
	})
}

func checkAssertString(t *testing.T, result, expect string) {
	t.Helper()

	if result != expect {
		t.Errorf("Wrong values, expect -> %s\n receive -> %s\n", expect, result)
	}
}

func checkAssertArray(t *testing.T, result, expect []string) {
	t.Helper()

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Wrong values, expect -> %s\n receive -> %s\n", expect, result)
	}
}
