package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("Welcome", 2))
}
func convert(s string, numRows int) string {

	for i := 0; i <= numRows; i++ {
		
		fmt.Println(string(s[i]), "\n")
	}

	return ""
}
