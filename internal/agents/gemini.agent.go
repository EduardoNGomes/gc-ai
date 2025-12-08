package agents

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/config"
	"google.golang.org/genai"
)

type GeminiAgent struct {
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

	diff, err := GetDiff()

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

func NewGeminiAgent() *GeminiAgent {
	return &GeminiAgent{}
}
