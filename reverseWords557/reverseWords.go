package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	arr := strings.Split(s, " ")
	arrTmp := []string{}
	for _, word := range arr {
		s = ""
		for i := len(word) - 1; i >= 0; i-- {
			s += string(word[i])
		}
		arrTmp = append(arrTmp, s)
	}
	

	return strings.Join(arrTmp, " ")
}

func main() {
	fmt.Println(reverseWords("set get let"))
}
