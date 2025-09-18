package agents

import (
	"context"
	"fmt"

	"github.com/eduardongomes/gcai/internal/config"
	"google.golang.org/genai"
)

type GeminiAgent struct{}

func (agent *GeminiAgent) GetCommit(config config.ConfigMethods, diff string) (string, error) {
	key := config.GetGeminiKey()

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  key,
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return "", fmt.Errorf("Error on get context: %w", err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(Prompt+diff),
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("error on generate content %w", err)
	}

	return result.Text(), nil
}

func (agent *GeminiAgent) Edit(v string) string {
	return ""
}

func NewGeminiAgent() *GeminiAgent {
	return &GeminiAgent{}
}
