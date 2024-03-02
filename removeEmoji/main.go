package main

import (
	"fmt"
	"regexp"
)

func removeEmojis(input string) string {
	emojiPattern := regexp.MustCompile(`[\x{1F600}-\x{1F64F}\x{1F300}-\x{1F5FF}\x{1F680}-\x{1F6FF}\x{2600}-\x{26FF}\x{2700}-\x{27BF}\x{2B50}\x{2B06}\x{2934}\x{2935}\x{200D}\x{23CF}\x{23E9}-\x{23F3}\x{231A}-\x{231B}\x{23E9}-\x{23F3}\x{23F8}-\x{23FA}]+`)
	cleanedText := emojiPattern.ReplaceAllString(input, "")

	return cleanedText
}

func main() {
	textWithEmojis := "🎉 Акция 3+"
	textWithoutEmojis := removeEmojis(textWithEmojis)

	fmt.Println("Original text:", textWithEmojis)
	fmt.Println("Text without emojis:", textWithoutEmojis)
}
