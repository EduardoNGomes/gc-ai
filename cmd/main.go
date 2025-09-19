package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/eduardongomes/gcai/internal/agents"
	"github.com/eduardongomes/gcai/internal/cli"
	"github.com/eduardongomes/gcai/internal/config"
)

func main() {
	openConfig := flag.Bool("config", false, "Enable new config keys")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	conf := config.NewConfig()
	cli := cli.NewCLI()
	geminiAgent := agents.NewGeminiAgent()

	cli.Run(*openConfig, conf, geminiAgent, reader)
}
