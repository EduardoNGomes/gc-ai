package menu

import (
	"github.com/manifoldco/promptui"
)

type MenuEditable interface {
	Run(label, defaultValue string) (string, error)
}

type ProdMenuEditable struct{}

func (m *ProdMenuEditable) Run(label, defaultValue string) (string, error) {
	promptConfirm := promptui.Prompt{
		Label:     label,
		Default:   defaultValue,
		AllowEdit: true,
	}

	line, err := promptConfirm.Run()
	if err != nil {
		return "", err
	}

	return line, nil
}

func NewProdMenuEditable() *ProdMenuEditable {
	return &ProdMenuEditable{}
}
