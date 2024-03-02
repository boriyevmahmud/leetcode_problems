package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(sortVowels("lEetcOde"))
}

var vowelsMap = map[rune]bool{
	'a': true,
	'e': true,
	'u': true,
	'o': true,
	'i': true,
	'A': true,
	'E': true,
	'U': true,
	'O': true,
	'I': true,
}

func sortVowels(s string) string {
	var vowel = make([]string, 0)
	for _, v := range s {
		if vowelsMap[(v)] {
			vowel = append(vowel, string(v))
		}
	}

	sort.Slice(vowel, func(i int, j int) bool {
		return vowel[i] < vowel[j]
	})
	cnt := 0
	resp := []rune(s)
	for i, v := range s {
		if vowelsMap[(v)] {
			resp[i] = []rune(vowel[cnt])[0]
			cnt++
		}
	}

	return string(resp)
}
