package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(fractionToDecimal(4, 333))

}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	sign := ""
	if (numerator < 0 && denominator > 0) || (numerator > 0 && denominator < 0) {
		sign = "-"
	}

	numerator = abs(numerator)
	denominator = abs(denominator)

	result := sign + strconv.Itoa(numerator/denominator)

	fraction := numerator % denominator
	if fraction == 0 {
		return result
	}
	result += "."
	remainderMap := make(map[int]int)

	for fraction != 0 {
		if pos, ok := remainderMap[fraction]; ok {
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}
		remainderMap[fraction] = len(result)
		fraction *= 10
		result += strconv.Itoa(fraction / denominator)
		fraction %= denominator
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
