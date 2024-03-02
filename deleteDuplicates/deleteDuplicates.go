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
	a := deleteDuplicates(&exm)
	for a != nil {
		fmt.Println(a.Val)
		a = a.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}



func deleteDuplicates(head *ListNode) *ListNode {
    res := head
    for head != nil {
        for head.Next != nil && head.Next.Val == head.Val {
            head.Next = head.Next.Next
        }
        head = head.Next
    }
    return res
}