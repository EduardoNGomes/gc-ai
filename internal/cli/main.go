package cli

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/agents"
	"github.com/eduardongomes/gcai/internal/config"
	c "github.com/eduardongomes/gcai/internal/config"
)

type CLI struct{}

type CLIMethdos interface {
	Run(openConfig bool, c c.ConfigMethods, a agents.AgentMethods, r io.Reader)
}

//go:embed .config.json
var embeddedConfig []byte

func (cli *CLI) Run(openConfig bool, c config.ConfigMethods, a agents.AgentMethods, reader io.Reader) {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	configPath := filepath.Join(exeDir, ".config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.WriteFile(configPath, embeddedConfig, 0644); err != nil {
			log.Fatal(err)
		}
	}

	if err := c.LoadEnvs(configPath); err != nil {
		log.Fatal(err)
	}

	r := c.IsEmpty()

	if r == true || openConfig {
		c.ConfigKey(reader)
	}

	msg, err := a.GetCommit(c)

	if err != nil {
		if errors.Is(err, errs.EmptyKeyError) {
			fmt.Println(err)
			return
		}
		log.Fatal(err)
	}

	if err := a.MakeCommit(msg); err != nil {
		log.Fatal(err)
	}
}

func NewCLI() *CLI {
	return &CLI{}
}
