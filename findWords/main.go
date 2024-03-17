package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findWords([]string{"ads", "fd", "ere"}))

}

func findWords(words []string) []string {
	var firstRow = map[rune]bool{
		'a': true,
		's': true,
		'd': true,
		'f': true,
		'g': true,
		'h': true,
		'j': true,
		'k': true,
		'l': true,
	}
	var secondRow = map[rune]bool{
		'q': true,
		'w': true,
		'e': true,
		'r': true,
		't': true,
		'y': true,
		'u': true,
		'i': true,
		'o': true,
		'p': true,
	}
	var thirdRow = map[rune]bool{
		'z': true,
		'x': true,
		'c': true,
		'v': true,
		'b': true,
		'n': true,
		'm': true,
	}

	resp := make([]string, 0, len(words))

	for _, v := range words {
		v1 := strings.ToLower(v)
		shouldAppend := true
		letters := make(map[rune]bool)

		if firstRow[rune(v1[0])] {
			letters = firstRow
		} else if secondRow[rune(v1[0])] {
			letters = secondRow
		} else if thirdRow[rune(v1[0])] {
			letters = thirdRow
		}

		for _, c := range v1 {
			if !letters[c] {
				shouldAppend = false
				break
			}
		}

		if shouldAppend {
			resp = append(resp, v)
		}

	}

	return resp
}
