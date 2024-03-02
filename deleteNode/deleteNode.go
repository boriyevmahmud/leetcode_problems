package main

import "fmt"

func main() {
	// 1,2,6,3,4,5,6
	exm := ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  9,
					Next: &ListNode{},
				},
			},
		},
	}
	deleteNode(&exm)
	for exm.Next != nil {
		fmt.Println(exm.Val)
		exm = *exm.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {

	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
