package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))

}

func findRepeatedDnaSequences(s string) []string {
	sequenceMap := make(map[string]int)
	repeatead := []string{}

	for i := 0; i < len(s)-9; i++ {
		seq := s[i : i+10]
		sequenceMap[seq]++
		if sequenceMap[seq] == 2 {
			repeatead = append(repeatead, seq)
		}
	}

	return repeatead
}
