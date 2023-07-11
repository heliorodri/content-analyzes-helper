package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

const API_KEY = "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"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please, type the PDF name, with whole path, that will be used as dataset...\n")
	scanner.Scan()
	file := scanner.Text()

	fmt.Print("Ask your question...\n")
	scanner.Scan()
	question := scanner.Text()

	textDataset := convertPdfToString(file)

	gptConfig := GptConfig{
		client: *initClient(API_KEY),
		ctx:    context.Background(),
	}

	answer := generateAnswers(question, textDataset, gptConfig)

	fmt.Println(answer)
}
