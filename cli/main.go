package cli

import (
	"io"

	c "github.com/eduardongomes/gcai/config"
)

type CLI struct{}

func (cli *CLI) Run(config c.ConfigMethods, reader io.Reader) {
	config.LoadEnvs()

	r := config.IsEmpty()

	if r == true {
		config.ConfigKey(reader)
	}
}

func NewCLI() *CLI {
	return &CLI{}
}
