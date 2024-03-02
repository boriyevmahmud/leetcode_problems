package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "ff200cf8-9507-4681-9859-e4356bc584c0:d4b1658f-3271-4973-8591-98a82939a664:1000"
	a := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(string(a))
}
