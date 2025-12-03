package prompt

import (
	"testing"

	l "github.com/eduardongomes/gcai/internal/line-reader"
	linereader "github.com/eduardongomes/gcai/internal/line-reader"
	"github.com/eduardongomes/gcai/internal/providers"
)

func TestCustomPrompt(t *testing.T) {

	customIndroduction := "Custom Introduction"
	customStructure := "Custom Structure"
	customExamples := []string{"Custom Examples"}
	customRules := []string{"Custom Rules"}

	reader := func() (l.LineReader, error) {
		return &mockReader{
			Inputs: []string{customIndroduction, customStructure, "n", "n"},
		}, nil
	}

	menu := &mockMenu{
		OptionsToReturn: []providers.MenuReturnOption{
			{Result: ""},
		},
	}
	t.Run("[GetIntroduction] Should return custom introduction", func(t *testing.T) {

		r, err := NewCustomPrompt(CustomPromptDTO{
			Introduction: customIndroduction,
			NewReader:    reader,
			MenuAction:   menu,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetIntroduction()
		expected := customIndroduction
		checkAssertString(t, result, expected)
	})
	t.Run("[GetStructure] Should return custom structure", func(t *testing.T) {
		sut, err := NewCustomPrompt(CustomPromptDTO{Structure: customStructure, NewReader: reader,
			MenuAction: menu,
		})

		if err != nil {
			t.Error(err)
		}
		result := sut.GetStructure()
		expected := customStructure

		checkAssertString(t, result, expected)
	})

	t.Run("[GetRules] Should return custom rules", func(t *testing.T) {

		r, err := NewCustomPrompt(CustomPromptDTO{
			Rules:      customRules,
			NewReader:  reader,
			MenuAction: menu,
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
			Examples:   customExamples,
			NewReader:  reader,
			MenuAction: menu,
		})

		if err != nil {
			t.Error(err)
		}

		result := r.GetExamples()

		expected := customExamples

		checkAssertArray(t, result, expected)
	})

	t.Run("[NewCustomPrompt]", func(t *testing.T) {
		initialArray := []string{"Regra 1"}
		nomeDoCampo := "RULES"

		mockMenu := &mockMenu{
			OptionsToReturn: []providers.MenuReturnOption{
				{Result: "ADD"},
			},
		}

		readerExterno := &mockReader{
			Inputs: []string{"y"},
		}

		readerInterno := &mockReader{
			Inputs: []string{"My new rule", "n"},
		}

		dto := &CustomPromptDTO{
			MenuAction: mockMenu,
			NewReader: func() (linereader.LineReader, error) {
				return readerInterno, nil
			},
		}

		resultArr, err := dto.editArrOption(readerExterno, initialArray, nomeDoCampo)

		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(resultArr) != 2 {
			t.Fatalf("Expected 2 itens, receive %d", len(resultArr))
		}

		expectedNewItem := "My new rule"
		if resultArr[1] != expectedNewItem {
			t.Errorf("Expected '%s', receive '%s'", expectedNewItem, resultArr[1])
		}
	})

}
