package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	reversed := make([]rune,0, len(s))

	vowels := []rune{}
	vowelsMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	for i := len(s) - 1; i >= 0; i-- {
		if vowelsMap[rune(s[i])] {
			vowels = append(vowels, rune(s[i]))
		}
	}
	cnt := 0
	for _, letter := range s {
		if vowelsMap[letter] {
			reversed = append(reversed, rune(vowels[cnt]))
			cnt++
		} else {
			reversed = append(reversed, letter)
		}
	}

	return string(reversed)
}
