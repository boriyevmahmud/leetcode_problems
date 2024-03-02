package main

import "fmt"

func main() {
	exm := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	a := removeDuplicates(&exm)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicates(head *ListNode) *ListNode {
	sential := &ListNode{Next: head}
	pred := sential
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Next.Val == head.Val {
				head.Next = head.Next.Next
			}
			pred.Next = head.Next
		} else {
			pred = pred.Next
		}

		head = head.Next
	}
	return sential.Next
}
