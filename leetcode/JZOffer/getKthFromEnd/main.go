package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if k <= 0 {
			slow = slow.Next
		}
		fast = fast.Next
		k--
	}
	return slow
}

func main() {
	head := ListNode{
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
	res := getKthFromEnd(&head, 2)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
