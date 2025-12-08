package agents

import (
	"context"
	"fmt"

	"github.com/eduardongomes/gcai/errs"
	"github.com/eduardongomes/gcai/internal/config"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAIAgent struct {
}

func (agent *OpenAIAgent) GetCommit(config config.ConfigMethods) (string, error) {
	key := config.GetOpenAIKey()

	if len(key) == 0 {
		return "", errs.EmptyKeyError
	}

	client := openai.NewClient(
		option.WithAPIKey(config.GetOpenAIKey()),
	)

	diff, err := GetDiff()

	if err != nil {
		return "", fmt.Errorf("error on generate content %w", err)
	}

	prompt := fmt.Sprintf("%s\nDIFF:\n%s", config.GetPromptString(), diff)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-4.1-nano",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
	})

	if err != nil {
		panic(err.Error())
	}

	return resp.OutputText(), nil
}

func NewOpenAIAgent() *OpenAIAgent {
	return &OpenAIAgent{}
}
