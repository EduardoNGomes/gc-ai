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
	f "github.com/eduardongomes/gcai/internal/flags"
	"github.com/eduardongomes/gcai/internal/providers"
)

type CLI struct{}

type CLIMethdos interface {
	Run(f f.Flags, c c.ConfigMethods, a agents.AgentMethods, r io.Reader)
}

//go:embed .config.json
var embeddedConfig []byte

func (cli *CLI) Run(f f.Flags, c config.ConfigMethods, a agents.AgentMethods, reader io.Reader) {

	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	configDir := filepath.Join(home, ".config", "gcai")

	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatal(err)
	}

	configPath := filepath.Join(configDir, "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := os.WriteFile(configPath, embeddedConfig, 0644); err != nil {
			log.Fatal(err)
		}
	}

	if err := c.LoadEnvs(configPath); err != nil {
		log.Fatal(err)
	}

	menu := providers.NewShellMenu("Choose your agent")
	agentsOptions := providers.NewSelectAgent(menu)

	if c.IsEmpty() || f.OpenConfig {
		if err := c.ConfigKey(reader, agentsOptions); err != nil {
			log.Fatal(err)
		}
	}

	if f.OpenConfig {
		fmt.Print("✅ Key saved")
		return
	}

	if f.AlterEditConfig != nil {
		if err := c.SetAllowEdit(*f.AlterEditConfig, true); err != nil {
			log.Fatal(err)
		}
		fmt.Println("✅ Config changes applied successfully.")
		return
	}

	msg, err := a.GetCommit(c)

	if err != nil {
		switch {
		case errors.Is(err, errs.EmptyKeyError):
			fmt.Println(err)
			return
		case errors.Is(err, errs.EmptyDiffError):
			fmt.Println(err)
			return
		default:
			log.Fatal(err)
		}
	}

	if c.GetAllowEdit() || f.EditCommit {
		msg, err = a.Edit(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := a.MakeCommit(msg); err != nil {
		log.Fatal(err)
	}
}

func NewCLI() *CLI {
	return &CLI{}
}
