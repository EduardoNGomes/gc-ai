package menu

import (
	"github.com/manifoldco/promptui"
)

type MenuConfirm interface {
	Run(label string) error
}

type ProdMenuConfirm struct{}

func (m *ProdMenuConfirm) Run(label string) error {
	promptConfirm := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	if _, err := promptConfirm.Run(); err != nil {
		return err
	}
	return nil

}

func NewProdMenuConfirm() *ProdMenuConfirm {
	return &ProdMenuConfirm{}
}
