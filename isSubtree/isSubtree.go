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
	reqSub := TreeNode{
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
	}
	fmt.Println(isSubtree(&req, &reqSub))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	if root.Val == subRoot.Val && IsIdentical(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func IsIdentical(t, s *TreeNode) bool {
	if s == nil || t == nil {
		return s == t
	}
	return s.Val == t.Val && IsIdentical(t.Left, s.Left) && IsIdentical(t.Right, s.Right)
}

// func isSubtree(s *TreeNode, t *TreeNode) bool {
// 	if s == nil || t == nil {
// 		return s == t
// 	}
// 	if s.Val == t.Val && isIdentical(s, t) {
// 		return true
// 	}
// 	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
// }

// func isIdentical(s, t *TreeNode) bool {
// 	if s == nil || t == nil {
// 		return s == t
// 	}
// 	return s.Val == t.Val && isIdentical(s.Left, t.Left) && isIdentical(s.Right, t.Right)
// }
