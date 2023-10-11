package main

import "fmt"

func main() {
	rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}

/*
1 2 3     7 4 1
4 5 6 	  8 5 2
7 8 9     9 6 3
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			fmt.Println("i: ", i, "j: ", j, "n-j-1: ", n-j-1)
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
			fmt.Println(matrix)
		}
	}
	// fmt.Println(matrix)
}
