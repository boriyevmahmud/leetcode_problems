package main

import "fmt"

func main() {
	fmt.Println(findSolution(temmFunc, 3))
}

func temmFunc(a int, b int) int {
	return a + b
}

func findSolution(customFunction func(int, int) int, z int) [][]int {
	resp := make([][]int, 0)
	x, y := 1, 1000
	fmt.Println("x", x)
	for x <= 1000 && y >= 1 {
		numb := customFunction(x, y)
		if numb == z {
			resp = append(resp, []int{x, y})
			x++
		} else if numb < z {
			x++
		} else {
			y--
		}
	}
	return resp
}

// var result [][]int
// 	x, y := 1, 1000 // Start from the bottom-left corner

// 	for x <= 1000 && y >= 1 {
// 		value := customFunction.f(x, y)
// 		if value == z {
// 			result = append(result, []int{x, y})
// 			x++
// 		} else if value < z {
// 			x++
// 		} else {
// 			y--
// 		}
// 	}

// 	return result
