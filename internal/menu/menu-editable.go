package menu

import (
	"github.com/charmbracelet/huh"
)

type MenuEditable interface {
	Run(label, defaultValue string) (string, error)
}

type ProdMenuEditable struct{}

func (m *ProdMenuEditable) Run(label, defaultValue string) (string, error) {
	var newValue string = defaultValue

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title(label).
				Value(&newValue).
				Lines(5),
		),
	)

	err := form.Run()
	if err != nil {
		return "", err
	}

	return newValue, nil
}

func NewProdMenuEditable() *ProdMenuEditable {
	return &ProdMenuEditable{}
}
