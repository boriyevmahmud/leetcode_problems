package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	abig, bBig, sum := new(big.Int), new(big.Int), new(big.Int)

	abig.SetString(a, 2)
	bBig.SetString(b, 2)

	return sum.Add(abig, bBig).Text(2)
}
