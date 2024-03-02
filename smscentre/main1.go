package main

import "fmt"

func main() {
	elem := []int{4, 34, 32, 43, 5, 24, 234, 54}
	for i := 0; i < len(elem)-1; i++ {
		fmt.Println(elem[i])
	}
	fmt.Println(elem[len(elem)-1])
}
