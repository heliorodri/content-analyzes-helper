package main

import (
	"context"
	"fmt"
	"io/ioutil"
)

const API_KEY = "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"

func main() {
	file := "/Users/helio.rodrigues/Desktop/test.pdf"
	text := convertPdfToString(file)
	// text := convertFileToString(file)

	err := ioutil.WriteFile("/Users/helio.rodrigues/Desktop/output.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Print("done")

	gptConfig := GptConfig{
		client: *initClient(API_KEY),
		ctx:    context.Background(),
	}

	generateBaseContent(text, gptConfig)
	answer := generateAnswers("se vc tivesse que escolher 1 bebedouro, qual seria?", gptConfig)

	fmt.Println(answer)
}
