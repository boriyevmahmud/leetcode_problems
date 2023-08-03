package main

import "fmt"

func main() {

	fmt.Println(isValid("(){}[]"))
}
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
			fmt.Println("in here after", stack)

		}
		for i := range stack {
			fmt.Println(string(stack[i]))
		}
	}

	return len(stack) == 0
}
