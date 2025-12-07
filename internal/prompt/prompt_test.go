package prompt_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/eduardongomes/gcai/internal/mocks"

	p "github.com/eduardongomes/gcai/internal/prompt"
)

func TestPrompt(t *testing.T) {
	introduction := "introduction test"
	structure := "structure test"
	rule := []string{"Rule 1"}
	example := []string{"Example 1"}
	t.Run("[ConvertToJSON] Should convert to JSON", func(t *testing.T) {
		prompt, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Introduction: introduction,
			Structure:    structure,
			Rules:        rule,
			Examples:     example,
			IsModify:     false,
		})

		if err != nil {
			t.Errorf("Error on convert JSON Prompt\nErr -> %v", err)
		}

		result, err := p.ConvertToJSON(prompt)

		if err != nil {
			t.Errorf("Error on convert JSON Prompt\nErr -> %v", err)
		}

		expectedObj := p.PromptJSON{
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

		result := p.ConvertStringArrayPromptToString(testArr)
		expect := `- Test 1
- Test 2`

		if result != expect {
			t.Errorf("Receive -> %s\nExpected -> %s", result, expect)
		}
	})

	t.Run("[ConvertToPromptString] Should create only one prompt string", func(t *testing.T) {
		prompt, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Introduction: introduction,
			Structure:    structure,
			Rules:        rule,
			Examples:     example,
			IsModify:     false,
		})

		if err != nil {
			t.Error(err)
		}

		result := p.ConvertToPromptString(prompt)
		fmt.Print(result)
		expect := fmt.Sprintf("Introduction:\n%s\nRules:\n%s\nStrucute:%s\nExamples:\n%s", introduction, p.ConvertStringArrayPromptToString(rule), structure, p.ConvertStringArrayPromptToString(example))

		if result != expect {
			t.Errorf("Receive -> %s\nExpected -> %s", result, expect)
		}

	})

	t.Run("[NewMenuPromptOptions] should return DEFAULT optiont", func(t *testing.T) {
		stub := mocks.NewStubMenuSelector()
		stub.Choices = append(stub.Choices, 0)

		r, err := p.NewMenuPromptOptions(stub)

		if err != nil {
			t.Errorf("Error on select Default Option -> %v ", err)
		}

		if r != p.PromptType("DEFAULT") {
			t.Fatalf("expected DEFAULT, got %s", r)
		}
	})

	t.Run("[NewMenuPromptOptions] Should return ERROR", func(t *testing.T) {
		stub := mocks.NewStubMenuSelector()
		stub.Choices = append(stub.Choices, 3)

		_, err := p.NewMenuPromptOptions(stub)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

}
