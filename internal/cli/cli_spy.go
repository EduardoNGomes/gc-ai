package cli

import (
	"io"
	"log"

	c "github.com/eduardongomes/gcai/internal/config"
	f "github.com/eduardongomes/gcai/internal/flags"
	"github.com/eduardongomes/gcai/internal/providers"
)

type CLISpy struct {
	testPath string
}

func (cli *CLISpy) Run(f f.Flags, config c.ConfigMethods, reader io.Reader) {

	if err := config.LoadEnvs(cli.testPath); err != nil {
		log.Fatal(err)
	}

	agents := providers.NewSelectAgentSpy()
	if config.IsEmpty() || f.OpenConfig {
		config.ConfigKey(reader, agents)
	}

}

func NewCLISpy() *CLISpy {
	return &CLISpy{}
}
