package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 5))
}

func searchMatrix(matrix [][]int, target int) bool {

	for _, mt := range matrix {
		for _, num := range mt {

			if target == num {
				return true
			}
		}
	}

	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
//     row :=len(matrix)-1
//     col := 0
//     for row >=0 && col < len(matrix[0]) {
//         if matrix[row][col] > target{
//             row--
//         } else if matrix[row][col] < target{
//             col++
//         } else {
//             return true
//         }
//     }
//     return false
// }
