package prompt

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"

	l "github.com/eduardongomes/gcai/internal/line-reader"
	"github.com/eduardongomes/gcai/internal/providers"
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
			IsModify:     false,
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
			IsModify:     false,
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

	t.Run("[NewMenuPromptOptions] should return DEFAULT optiont", func(t *testing.T) {
		stub := &stubMenu{
			DisplayFn: func() (*providers.MenuReturnOption, error) {
				return &providers.MenuReturnOption{
					Position: 0,
					Result:   "DEFAULT",
				}, nil
			},
		}
		r, err := NewMenuPromptOptions(stub)

		if err != nil {
			t.Errorf("Error on select Default Option -> %v ", err)
		}

		if r != PromptType("DEFAULT") {
			t.Fatalf("expected DEFAULT, got %s", r)
		}
	})

	t.Run("[NewMenuPromptOptions] Should return ERROR", func(t *testing.T) {
		mock := &stubMenu{
			DisplayFn: func() (*providers.MenuReturnOption, error) {
				return nil, errors.New("some error")
			},
		}

		_, err := NewMenuPromptOptions(mock)

		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

}

type stubMenu struct {
	DisplayFn func() (*providers.MenuReturnOption, error)
}

func (m *stubMenu) AddItem(label, value string) {}

func (m *stubMenu) Display() (*providers.MenuReturnOption, error) {
	return m.DisplayFn()
}
