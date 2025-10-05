package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/eduardongomes/gcai/internal/agents"
	"github.com/eduardongomes/gcai/internal/cli"
	"github.com/eduardongomes/gcai/internal/config"
	"github.com/eduardongomes/gcai/internal/flags"
)

func main() {
	openConfig := flag.Bool("config", false, "Enable new config keys")
	editCommit := flag.Bool("e", false, "Edit Commit Message")

	alterEditConfig := flag.String("edit", "", "Alter Configuration to edit commits")

	flag.Parse()

	var alterConfig *bool

	if *alterEditConfig != "" {
		val := strings.ToLower(*alterEditConfig) == "true"
		alterConfig = &val
	} else {
		alterConfig = nil
	}

	reader := bufio.NewReader(os.Stdin)

	conf := config.NewConfig()
	cli := cli.NewCLI()
	geminiAgent := agents.NewGeminiAgent()

	flags := flags.Flags{
		OpenConfig:      *openConfig,
		EditCommit:      *editCommit,
		AlterEditConfig: alterConfig,
	}

	cli.Run(flags, conf, geminiAgent, reader)
}
