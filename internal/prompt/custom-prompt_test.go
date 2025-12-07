package prompt_test

import (
	"testing"

	"github.com/eduardongomes/gcai/internal/menu"
	"github.com/eduardongomes/gcai/internal/mocks"

	p "github.com/eduardongomes/gcai/internal/prompt"
)

func TestCustomPrompt(t *testing.T) {

	customIndroduction := "Custom Introduction"
	customStructure := "Custom Structure"
	customExamples := []string{"Custom Examples"}
	customRules := []string{"Custom Rules"}

	stubMenuActionCommon := mocks.NewStubMenuSelector()
	stubMenuActionCommon.Choices = append(stubMenuActionCommon.Choices, 0)
	stubMenuActionCommon.Items = append(stubMenuActionCommon.Items, menu.MenuSelectorReturnOption{Position: 0, Result: "test"})

	t.Run("[GetIntroduction] Should return custom introduction", func(t *testing.T) {

		r, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Introduction: customIndroduction,
			MenuSelector: stubMenuActionCommon,
			IsModify:     false,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetIntroduction()
		expected := customIndroduction
		checkAssertString(t, result, expected)
	})
	t.Run("[GetStructure] Should return custom structure", func(t *testing.T) {
		sut, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Structure:    customStructure,
			MenuSelector: stubMenuActionCommon,
			IsModify:     false,
		})

		if err != nil {
			t.Error(err)
		}
		result := sut.GetStructure()
		expected := customStructure

		checkAssertString(t, result, expected)
	})

	t.Run("[GetRules] Should return custom rules", func(t *testing.T) {

		r, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Rules:        customRules,
			MenuSelector: stubMenuActionCommon,
			IsModify:     false,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetRules()
		expected := customRules

		checkAssertArray(t, result, expected)
	})
	t.Run("[GetExamples] Should return custom examples", func(t *testing.T) {

		r, err := p.NewCustomPrompt(p.CustomPromptDTO{
			Examples:     customExamples,
			MenuSelector: stubMenuActionCommon,
			IsModify:     false,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetExamples()

		expected := customExamples

		checkAssertArray(t, result, expected)
	})

	t.Run("[EditSTROption]", func(t *testing.T) {

		newRule := "Rule 1 edited"
		stubEditable := mocks.NewStubMenuEditable()
		stubEditable.Value = newRule

		dto := &p.CustomPromptDTO{
			MenuEditable: stubEditable,
			IsModify:     true,
		}

		result, err := dto.EditSTROption("test", "test")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if result != newRule {
			t.Errorf("Expected %s, receive %s", newRule, result)
		}
	})

	t.Run("[EditOptions - ADD]", func(t *testing.T) {

		stubSelector := mocks.NewStubMenuSelector()
		stubSelector.Choices = append(stubSelector.Choices, 0)

		stubConfirm := mocks.NewStubMenuConfirm()
		stubConfirm.Values = append(stubConfirm.Values, true)
		stubConfirm.Values = append(stubConfirm.Values, false)

		newValue := "Test 3"
		stubWriter := mocks.NewStubMenuWriter()
		stubWriter.Value = newValue

		dto := &p.CustomPromptDTO{
			MenuSelector: stubSelector,
			MenuConfirm:  stubConfirm,
			MenuWriter:   stubWriter,
			IsModify:     true,
		}

		initialArray := []string{"Test 1", "Test 2"}

		resultArr, err := dto.EditArrOption(initialArray, "Test ADD")

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if len(resultArr) != 3 {
			t.Errorf("Expected 3 itens, receive %d", len(resultArr))
		}

		if resultArr[2] != newValue {
			t.Errorf("Expected %s , receive %s", newValue, resultArr[2])
		}
	})

	t.Run("[EditOptions - Edit]", func(t *testing.T) {
		initialArray := []string{"Rule 1", "Rule 2"}
		fieldName := "RULES"

		stubSelector := mocks.NewStubMenuSelector()
		stubSelector.Choices = append(stubSelector.Choices, 1)
		stubSelector.Choices = append(stubSelector.Choices, 0)

		stubConfirm := mocks.NewStubMenuConfirm()
		stubConfirm.Values = append(stubConfirm.Values, true)
		stubConfirm.Values = append(stubConfirm.Values, false)

		newRule := "Rule 1 edited"
		stubEditable := mocks.NewStubMenuEditable()
		stubEditable.Value = newRule

		dto := &p.CustomPromptDTO{
			MenuSelector: stubSelector,
			MenuConfirm:  stubConfirm,
			MenuEditable: stubEditable,
			IsModify:     true,
		}

		resultArr, err := dto.EditArrOption(initialArray, fieldName)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(resultArr) != 2 {
			t.Errorf("Expected 2 item, receive %d", len(resultArr))
		}

		if resultArr[0] != newRule {
			t.Errorf("Expeted %s Receive %s", newRule, resultArr[0])
		}

	})

	t.Run("[EditOptions - REMOVE]", func(t *testing.T) {
		initialArray := []string{"Test 1", "Test 2"}
		fieldName := "Test"

		stubSelector := mocks.NewStubMenuSelector()
		stubSelector.Choices = append(stubSelector.Choices, 2)
		stubSelector.Choices = append(stubSelector.Choices, 1)

		stubConfirm := mocks.NewStubMenuConfirm()
		stubConfirm.Values = append(stubConfirm.Values, true)
		stubConfirm.Values = append(stubConfirm.Values, false)

		dto := &p.CustomPromptDTO{
			MenuSelector: stubSelector,
			MenuConfirm:  stubConfirm,
			IsModify:     true,
		}

		resultArr, err := dto.EditArrOption(initialArray, fieldName)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(resultArr) != 1 {
			t.Errorf("Expected 1 item, receive %d", len(resultArr))
		}
	})
}
