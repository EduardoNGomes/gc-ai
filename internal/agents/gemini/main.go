package gemini

import (
	"context"

	a "github.com/eduardongomes/gcai/internal/agents"
	c "github.com/eduardongomes/gcai/internal/config"
	"google.golang.org/genai"
)

type GeminiAgent struct{}

func (agent *GeminiAgent) GetCommit(config *c.Config, diff string) (string, error) {
	key := config.GetGeminiKey()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(a.Prompt+diff),
		nil,
	)

	if err != nil {
		return "", err
	}

	return result.Text(), nil
}
