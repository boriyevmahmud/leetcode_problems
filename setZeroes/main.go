package main

func main() {
	setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
}

/*
1 1 1
1 0 1
1 1 1
*/

func setZeroes(matrix [][]int) {
	indexes := [][]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				indexes = append(indexes, []int{i, j})
			}
		}
	}

	for _, val := range indexes {
		for k := 0; k < len(matrix[0]); k++ {
			matrix[val[0]][k] = 0
		}
		for k := 0; k < len(matrix); k++ {
			matrix[k][val[1]] = 0
		}
	}

}
