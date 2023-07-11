package main

import (
	"context"
	"fmt"
)

const API_KEY = "API-KEY-HERE"

func main() {
	file := "/Users/helio.rodrigues/Desktop/test.pdf"
	text := convertPdfToString(file)

	gptConfig := GptConfig{
		client: *initClient(API_KEY),
		ctx:    context.Background(),
	}

	answer := generateAnswers("se vc tivesse que escolher 1 bebedouro qual seria?", text, gptConfig)

	fmt.Println(answer)
}
