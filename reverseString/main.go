package main

import "fmt"

func main() {

	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func reverseString(s []byte) {

	lenS := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[lenS] = s[lenS], s[i]
		lenS -= 1
	}
	fmt.Println("s: ", string(s))

}
