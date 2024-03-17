package main

import "fmt"

func main() {

	fmt.Println(letterCombinations("12"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var output []string

	return nil
}
