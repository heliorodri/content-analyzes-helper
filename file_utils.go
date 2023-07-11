package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/pdfcrowd/pdfcrowd-go"
)

func convertPdfToString(filePath string) string {
	// create the API client instance
	client := pdfcrowd.NewPdfToTextClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	// run the conversion and store the result into the "txt" variable
	txt, err := client.ConvertFile(filePath)

	// check for the conversion error
	if err != nil {
		handleError(err)
	}

	return removeEmptyLines(removeNonPrintable(string(txt)))
}

func removeEmptyLines(text string) string {
	// Create a regular expression pattern to match empty lines
	emptyLinePattern := regexp.MustCompile(`(?m)^\s*$`)

	// Replace empty lines with an empty string
	textWithoutEmptyLines := emptyLinePattern.ReplaceAllString(text, "")

	// Trim leading and trailing spaces from each line
	textWithoutEmptyLines = strings.TrimSpace(textWithoutEmptyLines)

	return strings.ReplaceAll(textWithoutEmptyLines, "\n", "")
}

func removeNonPrintable(text string) string {
	// Create a regular expression pattern to match non-printable characters
	nonPrintablePattern := regexp.MustCompile(`[^\x20-\x7E]`)

	// Replace non-printable characters with an empty string
	asciiText := nonPrintablePattern.ReplaceAllString(text, "")

	return asciiText
}

func handleError(err error) {
	// report the error
	why, ok := err.(pdfcrowd.Error)
	if ok {
		os.Stderr.WriteString(fmt.Sprintf("Pdfcrowd Error: %s\n", why))
	} else {
		os.Stderr.WriteString(fmt.Sprintf("Generic Error: %s\n", err))
	}

	// rethrow or handle the exception
	panic(err.Error())
}
