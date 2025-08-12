package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("double  spaces"))
}

func ReverseWords(str string) string {
	// Tokenize the sentence
	var words []string = strings.Split(str, " ")
	var reverses = []string{}

	// Get each word
	for _, word := range words {
		var reverse string

		// Get each character and make a new word that's reversed
		for _, char := range word {
			reverse = string(char) + reverse
		}

		reverses = append(reverses, reverse)
	}

	return strings.Join(reverses, " ") // reverse those words
}
