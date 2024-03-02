package main

import "fmt"

func main() {
	// 1,2,6,3,4,5,6
	exm := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	a := reverseList(&exm)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{}
	}
	var newList *ListNode

	for head != nil {
		temp := head.Next
		head.Next = newList
		newList = head
		head = temp
	}

	return newList
}
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var revHead *ListNode

// 	for head != nil {
// 		tmp := head.Next
// 		head.Next = revHead
// 		revHead = head
// 		head = tmp
// 	}
// 	return revHead
// }
