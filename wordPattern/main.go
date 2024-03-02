package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog cat ct dog"))
}

func wordPattern(pattern string, s string) bool {
	splitted := strings.Split(s, " ")
	if len(splitted) != len(pattern) {
		return false
	}

	patternMap := make(map[rune]string)
	usedWords := make(map[string]bool)

	for i, p := range pattern {
		if word, ok := patternMap[p]; ok {
			// Check if the stored word for the pattern character is the same as the current word
			if word != splitted[i] {
				return false
			}
		} else {
			// Check if the current word has already been used for a different pattern character
			if usedWords[splitted[i]] {
				return false
			}

			patternMap[p] = splitted[i]
			usedWords[splitted[i]] = true
		}
	}

	return true
}
