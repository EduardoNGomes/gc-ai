package cli

import (
	"io"
	"log"
	"path/filepath"
	"runtime"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/agents"
	"github.com/eduardongomes/gcai/internal/config"
	c "github.com/eduardongomes/gcai/internal/config"
)

type CLI struct{}

type CLIMethdos interface {
	Run(c c.ConfigMethods, a agents.AgentMethods, r io.Reader)
}

func (cli *CLI) Run(c config.ConfigMethods, a agents.AgentMethods, reader io.Reader) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal(errs.CannotOpenFileErr)
	}

	configPath := filepath.Join(filepath.Dir(filename), "../.config.json")

	if err := c.LoadEnvs(configPath); err != nil {
		log.Fatal(err)
	}

	r := c.IsEmpty()

	if r == true {
		c.ConfigKey(reader)
	}

	msg, err := a.GetCommit(c)

	if err != nil {
		log.Fatal(err)
	}

	if err := a.MakeCommit(msg); err != nil {
		log.Fatal(err)
	}
}

func NewCLI() *CLI {
	return &CLI{}
}
