package main

import "fmt"

func main() {
	// 1,2,6,3,4,5,6
	exm := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	a := removeElements(&exm, 6)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// Create a dummy node to serve as the new head of the result list
	dummy := &ListNode{Next: head}
	prev, current := dummy, head

	// Iterate through the original list
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return dummy.Next
}
