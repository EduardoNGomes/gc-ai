package prompt

import (
	"fmt"
	"io"
	"log"
	"slices"
	"strings"

	linereader "github.com/eduardongomes/gcai/internal/line-reader"
	"github.com/eduardongomes/gcai/internal/providers"
)

type CustomPrompt struct {
	introduction string
	structure    string
	rules        []string
	examples     []string
}

type CustomPromptDTO struct {
	Introduction string
	Structure    string
	Rules        []string
	Examples     []string
	NewReader    func() (linereader.LineReader, error)
	OutputWriter io.Writer
	MenuAction   providers.Menu
}

func (p *CustomPrompt) GetIntroduction() string {
	return p.introduction
}

func (p *CustomPrompt) setIntroduction(v string) {
	p.introduction = v
}

func (p *CustomPrompt) GetStructure() string {
	return p.structure
}

func (p *CustomPrompt) setStructure(v string) {
	p.structure = v
}

func (p *CustomPrompt) GetRules() []string {
	return p.rules
}

func (p *CustomPrompt) setRules(v []string) {
	p.rules = v
}

func (p *CustomPrompt) GetExamples() []string {
	return p.examples
}

func (p *CustomPrompt) setExamples(v []string) {
	p.rules = v
}

func NewCustomPrompt(v CustomPromptDTO) (*CustomPrompt, error) {
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

	rules, err := v.editArrOption(rl, v.Rules, "RULES")

	if err != nil {
		return nil, err
	}

	examples, err := v.editArrOption(rl, v.Examples, "EXAMPLES")

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

func (c *CustomPromptDTO) editArrOption(rl linereader.LineReader, arr []string, name string) ([]string, error) {
	fmt.Printf("Do you want change %s? (Y/N) ", name)

	shouldEdit, err := rl.Readline()

	if err != nil {
		return arr, fmt.Errorf("error reading line: %v", err)
	}

	if strings.ToLower(shouldEdit) != "y" && strings.ToLower(shouldEdit) != "true" {
		return arr, nil
	}

	m := fmt.Sprintf("%s:", name)

	fmt.Println(m)
	fmt.Print(convertStringArrayPromptToString(arr))

	running := true

	rl, err = c.NewReader()

	if err != nil {
		return arr, err

	}
	defer rl.Close()

	for running {
		opt, err := c.MenuAction.Display()
		if err != nil {
			return arr, err
		}

		switch opt.Result {
		case "ADD":
			{
				m := fmt.Sprintf("Write your new %s:", name)
				fmt.Println(m)
				rl.SetPrompt("> ")
				userNewRule, err := rl.Readline()

				if err != nil {
					return arr, err
				}

				arr = append(arr, userNewRule)
				fmt.Printf("%s:", name)
				fmt.Print(convertStringArrayPromptToString(arr))
				break
			}

		case "EDIT":
			{
				opt, err := c.MenuAction.Display()
				if err != nil {
					log.Fatal(err)
				}

				rl.SetPrompt("Edit: ")
				rl.WriteStdin([]byte(opt.Result))
				userInput, err := rl.Readline()

				if err != nil {
					return arr, err
				}

				arr[opt.Position] = userInput

				fmt.Printf("%s:", name)
				fmt.Print(convertStringArrayPromptToString(arr))
				break
			}

		case "REMOVE":
			{
				opt, err := c.MenuAction.Display()
				if err != nil {
					return arr, err
				}

				arr = slices.Delete(arr, opt.Position, opt.Position+1)
				fmt.Printf("%s:", name)
				fmt.Print(convertStringArrayPromptToString(arr))
				break
			}
		default:
			{
				return arr, err
			}

		}

		fmt.Print("Should continue? (y/n) ")
		userContinueChoice, err := rl.Readline()

		if err != nil {
			return arr, err
		}

		if strings.ToLower(userContinueChoice) != "y" && strings.ToLower(userContinueChoice) != "true" && strings.ToLower(userContinueChoice) != "yes" {
			running = false
		}
	}
	return arr, nil
}

func NewMenuAction() providers.Menu {

	title := "Select an option"
	menu := providers.NewMenuCustomPrompt(title)

	menu.AddItem("ADD", "ADD")
	menu.AddItem("EDIT", "ADD")
	menu.AddItem("REMOVE", "ADD")

	return menu
}

func NewMenuOptions(name string, arr []string) providers.Menu {

	menu := providers.NewMenuCustomPrompt(name)

	for _, v := range arr {
		menu.AddItem(v, v)
	}

	return menu
}
