package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(reverseWords("  hello world  "))

}

func reverseWords(s string) string {
	arr := strings.Fields(s)
	respWord := ""
	for i := len(arr) - 1; i >= 1; i-- {
		if arr[i] != " " {
			respWord += arr[i] + " "
		}
	}
	respWord += arr[0]

	return respWord
}
