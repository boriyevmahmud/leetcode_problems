package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	exm := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{},
				},
			},
		},
	}
	swapPairs(&exm)

}

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	nextNode := head.Next
	head.Next = swapPairs(head.Next.Next)
	nextNode.Next = head

	return nextNode
}
