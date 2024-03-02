package main

import (
	"fmt"

	ph "github.com/dongri/phonenumber"
)

func main() {

	phone := "+998990265560"
	fmt.Println(ph.Parse(phone, "UZB"))
}
