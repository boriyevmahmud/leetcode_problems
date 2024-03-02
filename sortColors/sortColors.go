package main

import "fmt"

func main() {

	sortColors([]int{0, 2, 0, 0, 2, 1})
}

type Tmp struct {
	Red   int
	White int
	Blue  int
}

func sortColors(nums []int) {
	tmp := Tmp{}

	for _, num := range nums {
		switch num {
		case 0:
			tmp.Red += 1
		case 1:
			tmp.White += 1
		case 2:
			tmp.Blue += 1
		}
	}

	j := 0
	fmt.Println(tmp)
	for i := 0; i < tmp.Red; i++ {
		nums[j] = 0
		j++
	}

	for i := 0; i < tmp.White; i++ {
		nums[j] = 1
		j++
	}
	for i := 0; i < tmp.Blue; i++ {
		nums[j] = 2
		j++
	}

}
