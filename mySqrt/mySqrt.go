package main

import (
	"math"
)
 
func main() {
	mySqrt(8)
}

func mySqrt(x int) int {

	return int(math.Floor(math.Sqrt(float64(x))))
}
