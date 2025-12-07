package menu

import "github.com/manifoldco/promptui"

type MenuWriter interface {
	Run(label string) (string, error)
}

type ProdMenuWriter struct{}

func (m *ProdMenuWriter) Run(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
	}
	line, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return line, nil
}

func NewProdMenuWriter() *ProdMenuWriter {
	return &ProdMenuWriter{}
}
