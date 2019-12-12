package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	res := frequency(text)
	for _, kv := range res {
		fmt.Println(kv.Key+" - ", kv.Value)
	}
}

type kv struct {
	Key   string
	Value int
}

func frequency(input string) []kv {
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
	return getMostPopular(words)
}

func getMostPopular(words map[string]int) []kv {
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	return ss[0:10]
}

func isDelimiter(r rune, delimiters []rune) bool {
	for _, delim := range delimiters {
		if r == delim {
			return true
		}
	}
	return false
}
