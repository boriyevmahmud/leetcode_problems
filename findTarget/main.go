package main

import "fmt"

func main() {
	req := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(findTarget(&req, 0))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {

	return dfs(root, k, map[int]struct{}{})
}

func dfs(root *TreeNode, k int, m map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}

	m[root.Val] = struct{}{}
	return dfs(root.Left, k, m) || dfs(root.Right, k, m)

}
