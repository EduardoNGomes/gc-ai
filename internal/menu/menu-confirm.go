package menu

import (
	"errors"

	"github.com/charmbracelet/huh"
)

type MenuConfirm interface {
	Run(label string) error
}

type ProdMenuConfirm struct{}

func (m *ProdMenuConfirm) Run(label string) error {

	var confirmed bool

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(label).
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed),
		),
	).Run()

	if err != nil {
		return err
	}

	if !confirmed {
		return errors.New("canceled")
	}
	return nil
}

func NewProdMenuConfirm() *ProdMenuConfirm {
	return &ProdMenuConfirm{}
}
