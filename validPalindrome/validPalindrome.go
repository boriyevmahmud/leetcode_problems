package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

func isPalindrome(s string) bool {
    // Remove non-alphanumeric characters and spaces
    var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
    s = nonAlphanumericRegex.ReplaceAllString(s, "")
    s = strings.ReplaceAll(s, " ", "")
    s = strings.ToLower(s)

    length := len(s)
    mid := length / 2

    for i := 0; i < mid; i++ {
        j := length - i - 1
        if s[i] != s[j] {
            return false
        }
    }

    return true
}