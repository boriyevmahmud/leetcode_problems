package main

import "fmt"

func main() {
	req := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(countNodes(&req))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root.Right == nil && root.Left == nil {
		return 1
	}

	return countNodesWithInt(root)
}
func countNodesWithInt(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + countNodesWithInt(node.Left) + countNodesWithInt(node.Right)
}

// func countNodesWithInt(node *TreeNode, count int) int {
// 	if node == nil {
// 		return 0
// 	}
// 	count += 1

// 	if node.Right == nil && node.Left == nil {
// 		return count
// 	}

// 	return countNodesWithInt(node.Left, count) + countNodesWithInt(node.Right, count)
// }
