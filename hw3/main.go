package main

import (
	"fmt"
)

func main() {
	fmt.Println(frequency("asss ,sldfa sldfa sldfa alsdfa;, asdfa asdfa; alfs"))
}

func frequency(input string) string {
	delimiters := []rune{' ', '.', ',', '!', '?', ':', ';', '(', ')', '"', '\'', '-'}
	words := map[string]int{}
	lastIndex := 0
	wordReading := false

	for i, r := range input {
		if isDelimiter(r, delimiters) {
			if wordReading {
				word := input[lastIndex:i]
				words[word]++
			}
			wordReading = false
		} else {
			if !wordReading {
				lastIndex = i
			}
			wordReading = true
		}
	}
	if wordReading {
		word := input[lastIndex:len(input)]
		words[word]++
	}

	return getMax(words)
}

func getMax(words map[string]int) string {
	max := 0
	maxKey := ""
	for key, val := range words {
		if val > max {
			max = val
			maxKey = key
		}
	}
	return maxKey
}

func isDelimiter(r rune, delimiters []rune) bool {
	for _, delim := range delimiters {
		if r == delim {
			return true
		}
	}
	return false
}
