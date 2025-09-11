package cli

import (
	"io"
	"log"

	c "github.com/eduardongomes/gcai/internal/config"
)

type CLISpy struct {
	testPath string
}

func (cli *CLISpy) Run(config c.ConfigMethods, reader io.Reader) {

	if err := config.LoadEnvs(cli.testPath); err != nil {
		log.Fatal(err)
	}

	r := config.IsEmpty()

	if r == true {
		config.ConfigKey(reader)
	}
}

func NewCLISpy() *CLISpy {
	return &CLISpy{}
}
