package agents

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/chzyer/readline"
	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/config"
	l "github.com/eduardongomes/gcai/internal/line-reader"
	"google.golang.org/genai"
)

type GeminiAgent struct {
	newReader func() (l.LineReader, error)
	wasEdit   bool
}

func (agent *GeminiAgent) GetCommit(config config.ConfigMethods) (string, error) {
	key := config.GetGeminiKey()

	if len(key) == 0 {
		return "", errs.EmptyKeyError
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return "", fmt.Errorf("Error on get context: %w", err)
	}

	diff, err := agent.GetDiff()

	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf("%s\nDIFF:\n%s", config.GetPromptString(), diff)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("error on generate content %w", err)
	}

	return result.Text(), nil
}

var execCommand = exec.Command

func (agent *GeminiAgent) GetDiff() (string, error) {

	diff := execCommand("git", "diff", "--cached")

	stdout, err := diff.Output()

	if err != nil {
		return "", fmt.Errorf("Error on get git diff: %w", err)
	}

	d := string(stdout)

	if len(d) == 0 {
		return "", errs.EmptyDiffError

	}

	return d, nil
}

func (agent *GeminiAgent) MakeCommit(msg string) error {

	r := execCommand("git", "commit", "-m", msg)

	if _, err := r.Output(); err != nil {
		return fmt.Errorf("Erro on make commit: %v", err)
	}

	if !agent.wasEdit {
		fmt.Println(msg)
	}

	agent.wasEdit = false

	return nil
}

func (agent *GeminiAgent) Edit(msg string) (string, error) {
	rl, err := agent.newReader()

	if err != nil {
		return "", fmt.Errorf("error creating reader: %v", err)
	}
	defer rl.Close()

	fmt.Println("(Gemini)Please edit your commit message below, or press Enter to keep it unchanged:")

	rl.WriteStdin([]byte(msg))
	line, err := rl.Readline()
	if err != nil {
		return "", fmt.Errorf("error reading line: %v", err)
	}

	agent.wasEdit = true

	return line, nil
}

func NewGeminiAgent() *GeminiAgent {
	return &GeminiAgent{
		newReader: func() (l.LineReader, error) {
			return readline.New("")
		},
		wasEdit: false,
	}
}
