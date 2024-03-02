package main

import "fmt"

type St1 struct{}
type St2 struct{}

func main() {
	fmt.Println(St1{} == St1{})
}
