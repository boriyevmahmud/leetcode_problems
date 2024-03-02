package main

import "fmt"

func main() {
	fmt.Println(compress([]byte{61, 62, 62, 62, 63, 63, 63}))
}

func compress(chars []byte) int {
	str := ""
	stringsMap := make(map[byte]int)
	for _, v := range chars {
		stringsMap[v]++
	}
	for k, v := range stringsMap {
		if v == 1 {
			str += fmt.Sprintf("%v", string(k))
		} else {
			str += fmt.Sprintf("%v%v", string(k), v)
		}
	}
	for i, char := range chars {
		if v, ok := stringsMap[char]; ok {
			if v!=1 {
				
			}
		}
	}
	fmt.Println("str: ", str)
	return len(str)
}
