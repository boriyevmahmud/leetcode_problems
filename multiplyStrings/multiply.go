package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println(multiply("498828660196", "840477629533"))
}

func multiply(num1 string, num2 string) string {

	bigInt1, success1 := new(big.Int).SetString(num1, 10)
	bigInt2, success2 := new(big.Int).SetString(num2, 10)

	// Check if the conversions were successful
	if !success1 || !success2 {
		return "Invalid input"
	}

	// Multiply the big integers
	result := new(big.Int).Mul(bigInt1, bigInt2)

	return result.String()
}
