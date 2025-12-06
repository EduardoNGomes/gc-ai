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

type CLI struct {
	geminiAgent agents.AgentMethods
	openaiAgent agents.AgentMethods
}

type CLIMethdos interface {
	Run(f f.Flags, c c.ConfigMethods, r io.Reader)
}

//go:embed .config.json
var embeddedConfig []byte

func (cli *CLI) Run(f f.Flags, c config.ConfigMethods, reader io.Reader) {

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
		if err := c.Config(reader, agentsOptions, os.Stdout); err != nil {
			log.Fatal(err)
		}
	}

	if f.OpenConfig {
		fmt.Println("✅ Key saved")
		return
	}

	if f.AlterEditConfig != nil {

		c.RewriteConfig(config.RewriteConfigOptions{
			PromptType: nil,
			Prompt:     nil,
			AllowEdit:  nil,
			Agent:      nil,
		})

		if err := c.RewriteConfig(config.RewriteConfigOptions{
			PromptType: nil,
			Prompt:     nil,
			AllowEdit:  f.AlterEditConfig,
			Agent:      nil,
		}); err != nil {
			log.Fatal(err)
		}
		fmt.Println("✅ Config changes applied successfully.")
		return
	}

	if f.AlterAgent != nil {
		switch *f.AlterAgent {
		case "gemini":
			agent := providers.GEMINI
			{
				if err :=
					c.RewriteConfig(config.RewriteConfigOptions{
						PromptType: nil,
						Prompt:     nil,
						AllowEdit:  nil,
						Agent:      &agent,
					}); err != nil {
					log.Fatal(err)
				}
			}
		case "openai":
			agent := providers.OPEN_AI
			{
				if err := c.RewriteConfig(config.RewriteConfigOptions{
					PromptType: nil,
					Prompt:     nil,
					AllowEdit:  nil,
					Agent:      &agent,
				}); err != nil {
					log.Fatal(err)
				}
			}
		}
		fmt.Println("✅ New agent seleted")
		return
	}

	var agent agents.AgentMethods

	switch c.GetAgent() {
	case providers.GEMINI:
		{
			agent = cli.geminiAgent
		}
	case providers.OPEN_AI:
		{
			agent = cli.openaiAgent
		}
	default:
		{
			log.Fatal(errs.InvalidAgentSelected)
		}
	}

	msg, err := agent.GetCommit(c)

	if f.OnlyShowCommit {
		fmt.Println(msg)
		return
	}

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
		msg, err = agent.Edit(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := agent.MakeCommit(msg); err != nil {
		log.Fatal(err)
	}
}

func NewCLI(cfg struct {
	Gemini agents.AgentMethods
	OpenAI agents.AgentMethods
}) *CLI {
	return &CLI{
		geminiAgent: cfg.Gemini,
		openaiAgent: cfg.OpenAI,
	}
}
