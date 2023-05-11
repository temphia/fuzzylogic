package fuzzymod

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type FuzzyMod struct {
	client *openai.Client
}

func New(apiKey string) *FuzzyMod {

	client := openai.NewClient(apiKey)

	return &FuzzyMod{
		client: client,
	}
}

func (fm *FuzzyMod) ChatCompletion(opts *openai.ChatCompletionRequest) (any, error) {

	resp, err := fm.client.CreateChatCompletion(context.Background(), *opts)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
