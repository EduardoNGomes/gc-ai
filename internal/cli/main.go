package cli

import (
	"io"
	"log"
	"path/filepath"
	"runtime"

	"github.com/eduardongomes/gcai/errs"
	c "github.com/eduardongomes/gcai/internal/config"
)

type CLI struct{}

type CLIMethdos interface {
	Run(c c.ConfigMethods, r io.Reader)
}

func (cli *CLI) Run(config c.ConfigMethods, reader io.Reader) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal(errs.CannotOpenFileErr)
	}

	configPath := filepath.Join(filepath.Dir(filename), "../.config.json")

	if err := config.LoadEnvs(configPath); err != nil {
		log.Fatal(err)
	}

	r := config.IsEmpty()

	if r == true {
		config.ConfigKey(reader)
	}
}

func NewCLI() *CLI {
	return &CLI{}
}
