package main

import "fmt"

func main() {
	req := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(binaryTreePaths(&req))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {

	return binaryTreePathsWithString(root, "")
}

func binaryTreePathsWithString(node *TreeNode, str string) []string {
	if node == nil {
		return []string{}
	}

	if node.Right == nil && node.Left == nil {
		str += fmt.Sprintf("%v", node.Val)
		return []string{str}
	}

	str += fmt.Sprintf("%v->", node.Val)

	return append(
		binaryTreePathsWithString(node.Left, str),
		binaryTreePathsWithString(node.Right, str)...,
	)
}
