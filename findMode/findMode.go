package main

import "fmt"

func main() {
	req := TreeNode{
		Val:  3,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 2,
				},
				Left: &TreeNode{
					Val: 2,
				},
			},
		},
	}
	fmt.Println(findMode(&req))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	countsMap := make(map[int]int)
	findModeWithMap(root, countsMap)
	return findModesFromMap(countsMap)
}

func findModeWithMap(node *TreeNode, countsMap map[int]int) {
	if node != nil {
		countsMap[node.Val] += 1
		findModeWithMap(node.Left, countsMap)
		findModeWithMap(node.Right, countsMap)
	}
}

func findModesFromMap(data map[int]int) []int {
	var modes []int
	var maxCount int

	for _, count := range data {
		if count > maxCount {
			maxCount = count
		}
	}

	for number, count := range data {
		if count == maxCount {
			modes = append(modes, number)
		}
	}

	return modes
}
