package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("adf   "))
}

func lengthOfLastWord(s string) int {
	if !strings.Contains(s, " ") {
		return len(s)
	}
	counter := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			counter++
		}
		if (counter > 1 && s[i] == ' ') || i == 0 {
			return counter
		}
	}

	return 1
}
