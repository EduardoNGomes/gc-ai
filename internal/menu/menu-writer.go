package menu

import "github.com/charmbracelet/huh"

type MenuWriter interface {
	Run(label string) (string, error)
}

type ProdMenuWriter struct{}

func (m *ProdMenuWriter) Run(label string) (string, error) {
	var line string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(label).
				Value(&line),
		),
	).Run()

	if err != nil {
		return "", err
	}

	return line, nil
}

func NewProdMenuWriter() *ProdMenuWriter {
	return &ProdMenuWriter{}
}
