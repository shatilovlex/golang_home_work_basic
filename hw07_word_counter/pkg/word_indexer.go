package pkg

import (
	"regexp"
	"strings"
)

func countWords(input string) map[string]int {
	words := make(map[string]int)
	wordOnly := punctCleaner(input)
	splitWords := strings.Fields(wordOnly)

	for _, word := range splitWords {
		words[word]++
	}

	return words
}

func punctCleaner(input string) string {
	return regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(input, "")
}
