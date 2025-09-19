package main

import (
	"bufio"
	"os"

	"github.com/eduardongomes/gcai/internal/agents"
	"github.com/eduardongomes/gcai/internal/cli"
	"github.com/eduardongomes/gcai/internal/config"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conf := config.NewConfig()
	cli := cli.NewCLI()
	geminiAgent := agents.NewGeminiAgent()

	cli.Run(conf, geminiAgent, reader)

	//	r := exec.Command("git", "diff", "--cached")
	//
	//	stdout, err := r.Output()
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	input := string(stdout)
	//
	//	ctx := context.Background()
	//	client, err := genai.NewClient(ctx, &genai.ClientConfig{
	//		APIKey:  "AIzaSyDModzmImT0_iZhzNHdUUVPs3d2LZrVO_U",
	//		Backend: genai.BackendGeminiAPI,
	//	})
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	prompt := `You are a senior software engineer. Analyze the git diff provided below and output **only** the single-line conventional commit message it represents.
	//
	// - The message must follow the Conventional Commits specification: 'type(scope): description'.
	//   - Choose the most appropriate type from: 'feat', 'fix', 'refactor', 'chore', 'docs', 'style', 'perf', 'test'.
	//
	// `
	//
	//	//prompt := "You are a senior software engineer. Analyze the given git diff and output only a single-line, professional conventional commit message following the Conventional Commits specification (type, scope if applicable, and short description), use feat,refator,fix,chore when needed "
	//
	//	result, err := client.Models.GenerateContent(
	//		ctx,
	//		"gemini-2.5-flash-lite",
	//		genai.Text(prompt+input),
	//		nil,
	//	)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//
	//	}
	//
	//	fmt.Println(result.Text())
	//
	//	r2 := exec.Command("git", "commit", "-m", result.Text())
	//
	//	if _, err := r2.Output(); err != nil {
	//		log.Fatal(err)
	//	}
}
