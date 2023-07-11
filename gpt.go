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

func generateBaseContent(context string, config GptConfig) openai.ChatCompletionResponse {
	messages := openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: context,
	}

	request := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo16K,
		Messages: []openai.ChatCompletionMessage{messages},
	}

	completion, err := config.client.CreateChatCompletion(config.ctx, request)

	if err != nil {
		panic(err)
	}

	return completion
}

func generateAnswers(input string, config GptConfig) string {
	request := openai.CompletionRequest{
		Model:  openai.GPT3Dot5Turbo16K,
		Prompt: []string{input},
	}

	completion, err := config.client.CreateCompletion(config.ctx, request)

	if err != nil {
		panic(err)
	}

	return completion.Choices[0].Text
}
