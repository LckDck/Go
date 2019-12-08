package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to expand: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Expanded : " + expand(text))
}

func expand(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	var amount int = 0
	var pow int = 0
	var lastSymbol rune = runes[0]
	const escape = 92
	var nextEscaped = false

	var b strings.Builder

	for i, r := range runes {
		if i == 0 {
			continue
		}
		if r == escape {
			writeToBuilder(&b, lastSymbol, amount, nextEscaped)
			nextEscaped = true
			continue
		}
		if nextEscaped {
			lastSymbol = r
			nextEscaped = false
			continue
		}
		if unicode.IsDigit(r) {
			amount = amount*10 + (int(r) - '0')
			pow++
			continue
		} else {
			writeToBuilder(&b, lastSymbol, amount, nextEscaped)
			lastSymbol = r
			amount = 0
			pow = 0
		}
	}
	writeToBuilder(&b, lastSymbol, amount, nextEscaped)
	return b.String()
}

func writeToBuilder(b *strings.Builder, s rune, amount int, allowAll bool) {
	// if !unicode.IsDigit(s) || allowAll {
	b.WriteRune(s)
	for j := 0; j < amount-1; j++ {
		b.WriteRune(s)
	}
	//}
}
