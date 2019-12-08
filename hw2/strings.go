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

	var b strings.Builder

	for i, r := range runes {
		if i == 0 {
			continue
		}
		if unicode.IsDigit(r) {
			amount = amount*10 + (int(r) - '0')
			pow++
			continue
		} else {
			writeToBuilder(&b, lastSymbol, amount)
			lastSymbol = r
			amount = 0
			pow = 0
		}
	}
	writeToBuilder(&b, lastSymbol, amount)
	return b.String()
}

func writeToBuilder(b *strings.Builder, s rune, amount int) {
	if !unicode.IsDigit(s) {
		b.WriteRune(s)
		for j := 0; j < amount-1; j++ {
			b.WriteRune(s)
		}
	}
}
