package agents

import (
	"context"
	"fmt"

	"github.com/chzyer/readline"
	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/config"
	l "github.com/eduardongomes/gcai/internal/line-reader"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAIAgent struct {
	newReader func() (l.LineReader, error)
	wasEdit   bool
}

func (agent *OpenAIAgent) GetCommit(config config.ConfigMethods) (string, error) {
	key := config.GetOpenAIKey()

	if len(key) == 0 {
		return "", errs.EmptyKeyError
	}

	client := openai.NewClient(
		option.WithAPIKey(config.GetOpenAIKey()),
	)

	diff, err := agent.GetDiff()

	if err != nil {
		return "", fmt.Errorf("error on generate content %w", err)
	}

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-4.1-nano",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(Prompt + diff)},
	})

	if err != nil {
		panic(err.Error())
	}

	return resp.OutputText(), nil
}

func (agent *OpenAIAgent) GetDiff() (string, error) {

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

func (agent *OpenAIAgent) MakeCommit(msg string) error {

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

func (agent *OpenAIAgent) Edit(msg string) (string, error) {
	rl, err := agent.newReader()

	if err != nil {
		return "", fmt.Errorf("error creating reader: %v", err)
	}
	defer rl.Close()

	fmt.Println("Please edit your commit message below, or press Enter to keep it unchanged:")

	rl.WriteStdin([]byte(msg))
	line, err := rl.Readline()
	if err != nil {
		return "", fmt.Errorf("error reading line: %v", err)
	}

	agent.wasEdit = true

	return line, nil
}

func NewOpenAIAgent() *OpenAIAgent {
	return &OpenAIAgent{
		newReader: func() (l.LineReader, error) {
			return readline.New("")
		},
		wasEdit: false,
	}
}
