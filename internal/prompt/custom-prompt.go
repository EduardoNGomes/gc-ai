package prompt

import (
	"fmt"
	"io"
	"slices"
	"strings"

	linereader "github.com/eduardongomes/gcai/internal/line-reader"
	m "github.com/eduardongomes/gcai/internal/menu"
)

type CustomPrompt struct {
	introduction string
	structure    string
	rules        []string
	examples     []string
}

type CustomPromptJSON struct {
	Introduction string   `json:"introduction"`
	Structure    string   `json:"structure"`
	Examples     []string `json:"exemples"`
	Rules        []string `json:"Rules"`
}

type CustomPromptDTO struct {
	Introduction string
	Structure    string
	Rules        []string
	Examples     []string
	NewReader    func() (linereader.LineReader, error)
	OutputWriter io.Writer
	MenuSelector m.MenuSelector
	MenuConfirm  m.MenuConfirm
	MenuWriter   m.MenuWriter
	MenuEditable m.MenuEditable
	IsModify     bool
}

func (p *CustomPrompt) GetIntroduction() string {
	return p.introduction
}

func (p *CustomPrompt) GetStructure() string {
	return p.structure
}

func (p *CustomPrompt) GetRules() []string {
	return p.rules
}

func (p *CustomPrompt) GetExamples() []string {
	return p.examples
}

func NewCustomPrompt(v CustomPromptDTO) (Prompt, error) {
	if !v.IsModify {
		return &CustomPrompt{
			introduction: v.Introduction,
			structure:    v.Structure,
			rules:        v.Rules,
			examples:     v.Examples,
		}, nil
	}

	rl, err := v.NewReader()

	if err != nil {
		return nil, fmt.Errorf("error creating reader: %v", err)
	}

	defer rl.Close()

	introduction, err := v.editSTROption(rl, "Introduction", v.Introduction)

	if err != nil {
		return nil, err
	}

	structure, err := v.editSTROption(rl, "Structure", v.Structure)

	if err != nil {
		return nil, err
	}

	rules, err := v.editArrOption(v.Rules, "RULES")

	if err != nil {
		return nil, err
	}

	examples, err := v.editArrOption(v.Examples, "EXAMPLES")

	if err != nil {
		return nil, err
	}

	return &CustomPrompt{
		introduction: introduction,
		structure:    structure,
		rules:        rules,
		examples:     examples,
	}, nil
}

func (c *CustomPromptDTO) editSTROption(rl linereader.LineReader, name, value string) (string, error) {

	m := fmt.Sprintf("Write your prompt %s or press ENTER to keep it unchanged:", name)

	fmt.Println(m)

	oldValue := value

	rl.WriteStdin([]byte(oldValue))
	newValue, err := rl.Readline()

	if err != nil {
		return oldValue, fmt.Errorf("error reading line: %v", err)
	}

	return newValue, nil
}

func (c *CustomPromptDTO) editArrOption(arr []string, name string) ([]string, error) {
	msg := fmt.Sprintf("Do you want change %s", name)
	if err := c.MenuConfirm.Run(msg); err != nil {
		return arr, nil
	}

	m := fmt.Sprintf("%s:", name)

	fmt.Println(m)
	fmt.Print(convertStringArrayPromptToString(arr))

	running := true

	actionOptions := []string{"ADD", "EDIT", "REMOVE"}

	for running {
		opt, err := c.MenuSelector.Run("Choose an action:", actionOptions)
		if err != nil {
			return arr, err
		}

		switch opt.Result {
		case "ADD":
			{
				msg := fmt.Sprintf("Write your new %s:", name)

				newRule, err := c.MenuWriter.Run(msg)

				if err != nil {
					return arr, err
				}

				if strings.TrimSpace(newRule) != "" {
					arr = append(arr, newRule)
				}

				fmt.Printf("%s:\n", name)
				fmt.Println(convertStringArrayPromptToString(arr))
				break
			}

		case "EDIT":
			{
				opt, err := c.MenuSelector.Run("Chose one to edit:", arr)
				if err != nil {
					return arr, err
				}

				line, err := c.MenuEditable.Run("Edit:", opt.Result)

				if err != nil {
					return arr, err
				}

				arr[opt.Position] = line

				fmt.Printf("%s:\n", name)
				fmt.Println(convertStringArrayPromptToString(arr))
				break
			}

		case "REMOVE":
			{
				opt, err := c.MenuSelector.Run("Chose one to edit:", arr)
				if err != nil {
					return arr, err
				}

				arr = slices.Delete(arr, opt.Position, opt.Position+1)
				fmt.Printf("%s:\n", name)
				fmt.Println(convertStringArrayPromptToString(arr))
				break
			}
		default:
			{
				return arr, err
			}

		}

		if err := c.MenuConfirm.Run("Should continue? (y/n)"); err != nil {
			running = false
		}
	}
	return arr, nil
}
