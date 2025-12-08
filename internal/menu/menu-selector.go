package menu

import "github.com/charmbracelet/huh"

type MenuSelectorReturnOption struct {
	Position int
	Result   string
}

type MenuSelector interface {
	Run(label string, items []string) (*MenuSelectorReturnOption, error)
}

type ProdMenuSelector struct{}

func (m *ProdMenuSelector) Run(label string, items []string) (*MenuSelectorReturnOption, error) {
	var result string

	options := make([]huh.Option[string], len(items))
	for i, v := range items {
		options[i] = huh.NewOption(v, v)
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(label).
				Options(options...).
				Value(&result),
		),
	).Run()

	if err != nil {
		return nil, err
	}

	position := -1
	for i, v := range items {
		if v == result {
			position = i
			break
		}
	}

	return &MenuSelectorReturnOption{Position: position, Result: result}, nil

}

func NewProdMenuSelector() *ProdMenuSelector {
	return &ProdMenuSelector{}
}
