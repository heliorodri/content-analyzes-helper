package main

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type GptConfig struct {
	client openai.Client
	ctx    context.Context
}

func initClient(apikey string) *openai.Client {
	return openai.NewClient(apikey)
}

func generateAnswers(input string, context string, config GptConfig) string {
	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: context,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: input,
			},
		},
	}

	completion, err := config.client.CreateChatCompletion(config.ctx, request)

	if err != nil {
		panic(err)
	}

	return completion.Choices[0].Message.Content
}
