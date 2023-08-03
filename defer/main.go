package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	aes.NewCipher([]byte{0101, 0})
}
