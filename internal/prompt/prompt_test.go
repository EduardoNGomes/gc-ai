package prompt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	l "github.com/eduardongomes/gcai/internal/line-reader"
)

func TestPrompt(t *testing.T) {
	introduction := "introduction test"
	structure := "structure test"
	rule := []string{"Rule 1"}
	example := []string{"Example 1"}

	reader := func() (l.LineReader, error) {
		return &mockReader{
			Inputs: []string{introduction, structure, rule[0], example[0]},
		}, nil
	}

	t.Run("[ConvertToJSON] Should convert to JSON", func(t *testing.T) {
		p, err := NewCustomPrompt(CustomPromptDTO{
			Introduction: introduction,
			Structure:    structure,
			Rules:        rule,
			Examples:     example,
			NewReader:    reader,
		})

		if err != nil {
			t.Errorf("Error on convert JSON Prompt\nErr -> %v", err)
		}

		result, err := ConvertToJSON(p)

		if err != nil {
			t.Errorf("Error on convert JSON Prompt\nErr -> %v", err)
		}

		expectedObj := PromptJSON{
			Introduction: introduction,
			Rules:        rule,
			Structure:    structure,
			Examples:     example,
		}

		expect, _ := json.Marshal(expectedObj)

		if !reflect.DeepEqual(result, expect) {
			t.Errorf("\nReceive -> %s\nExpected -> %s", string(result), string(expect))
		}
		if !reflect.DeepEqual(result, expect) {
			t.Errorf("Receive -> %s\nExpected -> %s", result, expect)
		}
	})

	t.Run("[convertStringArrayPromptToString] Should join array,following determinate structure", func(t *testing.T) {

		testArr := []string{
			"Test 1",
			"Test 2",
		}

		result := convertStringArrayPromptToString(testArr)
		expect := `- Test 1
- Test 2`

		if result != expect {
			t.Errorf("Receive -> %s\nExpected -> %s", result, expect)
		}
	})

	t.Run("[ConvertToPromptString] Should create only one prompt string", func(t *testing.T) {
		p, err := NewCustomPrompt(CustomPromptDTO{
			Introduction: introduction,
			Structure:    structure,
			Rules:        rule,
			Examples:     example,
			NewReader:    reader,
		})

		if err != nil {
			t.Error(err)
		}

		result := ConvertToPromptString(p)
		fmt.Print(result)
		expect := fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", introduction, convertStringArrayPromptToString(rule), structure, convertStringArrayPromptToString(example))

		if result != expect {
			t.Errorf("Receive -> %s\nExpected -> %s", result, expect)
		}

	})

}
