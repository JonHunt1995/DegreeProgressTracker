package main

import "strings"

func sanitizeInputs(slicedInput string) []string {
	words := strings.Fields(slicedInput)
	sanitizedWords := []string{}

	for _, word := range words {
		word = strings.ToLower(word)
		sanitizedWords = append(sanitizedWords, word)
	}
	return sanitizedWords
}
