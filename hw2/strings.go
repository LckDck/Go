package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to expand: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Expanded : " + expand(text))
}

func expand(s string) string {
	const escape = 92
	runes := []rune(s)
	result := []rune{}
	escapeActive := false
	amount := 0

	for _, r := range runes {
		if escapeActive || (r != escape && !unicode.IsDigit(r)) {
			result = applyAmount(result, amount)
			result = append(result, r)
			amount = 0
			escapeActive = false
		} else if r == escape {
			escapeActive = true
		} else if unicode.IsDigit(r) {
			amount = amount*10 + (int(r) - '0')
		}
	}
	if amount > 0 {
		result = applyAmount(result, amount)
	}
	return string(result)
}

func applyAmount(result []rune, amount int) []rune {
	length := len(result)
	if length > 0 && amount > 0 {
		last := result[length-1]
		for j := 0; j < amount-1; j++ {
			result = append(result, last)
		}
	}
	return result
}
