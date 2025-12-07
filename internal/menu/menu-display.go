package providers

import (
	"github.com/manifoldco/promptui"
)

type MenuSelectorReturnOption struct {
	Position int
	Result   string
}

type MenuSelector interface {
	Run(label string, items []string) (*MenuSelectorReturnOption, error)
}

type ProdMenuSelector struct{}

func (m *ProdMenuSelector) Run(label string, items []string) (*MenuSelectorReturnOption, error) {
	prompt := promptui.Select{
		HideHelp: true,
		Label:    label,
		Items:    items,
		Size:     6,
	}

	i, result, err := prompt.Run()

	if err != nil {
		return nil, err
	}

	return &MenuSelectorReturnOption{Position: i, Result: result}, nil

}

func NewProdMenuSelector() *ProdMenuSelector {
	return &ProdMenuSelector{}
}
