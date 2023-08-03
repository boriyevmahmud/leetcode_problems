package main

import (
	"fmt"
	"strings"
)

func main() {
	comment := "Your comment with special characters !@#$%^&*()"

	// Add backslash before each special character
	escapedComment := strings.NewReplacer(
		"!", "\\!",
		"@", "\\@",
		"#", "\\#",
		"$", "\\$",
		"%", "\\%",
		"^", "\\^",
		"&", "\\&",
		"*", "\\*",
		"(", "\\(",
		")", "\\)",
	).Replace(comment)

	fmt.Println(escapedComment)
}
