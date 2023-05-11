package fuzzyreact

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatBot struct {
	system   string
	messages []openai.ChatCompletionMessage
	client   *openai.Client
}

func newChatBot(client *openai.Client, system string) *ChatBot {
	messages := []openai.ChatCompletionMessage{}
	if system != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: system,
		})
	}

	return &ChatBot{
		system:   system,
		messages: messages,
		client:   client,
	}

}

func (cb *ChatBot) execute(msg string) (string, error) {

	cb.messages = append(cb.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	result, err := cb.doExecute()
	if err != nil {
		return "", err
	}

	cb.messages = append(cb.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: result,
	})

	return result, nil

}

func (cb *ChatBot) doExecute() (string, error) {

	completion, err := cb.client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo0301,
			Messages: cb.messages,
		})

	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}

func (cb *ChatBot) addMessage(role, content string) {

	cb.messages = append(cb.messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	})

}
