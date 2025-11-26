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

	alterEditConfig := flag.String("edit", "", "Alter Configuration to edit commits(true/false)")
	alterAgentConfig := flag.String("agent", "", "Alter Agent to make commits(gemini/openai)")

	flag.Parse()

	var alterConfig *bool
	var alterAgent *string

	if *alterEditConfig != "" {
		val := strings.ToLower(*alterEditConfig) == "true"
		alterConfig = &val
	} else {
		alterConfig = nil
	}

	if *alterAgentConfig != "" {
		if strings.ToLower(*alterEditConfig) == "gemini" || strings.ToLower(*alterEditConfig) == "openai" {
			alterAgent = alterEditConfig
		}
	} else {
		alterAgent = nil
		alterConfig = nil
	}

	reader := bufio.NewReader(os.Stdin)

	conf := config.NewConfig()

	geminiAgent := agents.NewGeminiAgent()
	openAIAgent := agents.NewOpenAIAgent()

	cli := cli.NewCLI(struct {
		Gemini agents.AgentMethods
		OpenAI agents.AgentMethods
	}{
		Gemini: geminiAgent,
		OpenAI: openAIAgent,
	})

	flags := flags.Flags{
		OpenConfig:      *openConfig,
		EditCommit:      *editCommit,
		AlterEditConfig: alterConfig,
		AlterAgent:      alterAgent,
	}

	cli.Run(flags, conf, reader)
}
