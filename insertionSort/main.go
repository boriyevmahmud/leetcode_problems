package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11,43,1,2,4,32,4,3,234,4324,32,43,432,54,6,5,6,7,67,0,5,423,5,432,54,23,2,54,3,5,64,45,36,432,54,25,42,35,24,4,3,2,65,62,4,2,45,4,54,45,65,6,24,5,623,5,2435,24,35,243,524,35,423,5,4235,4,65,65,0,4,5,243,5,2435,24,54,62,55,6,7,8,87,54,4,34,23,4,3,45,3,2,3,4,4352,2}
	fmt.Println("Unsorted array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
