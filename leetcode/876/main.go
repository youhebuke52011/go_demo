package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
链表中间节点
思路：快慢指针，慢指针走一步，快指针走两步，当快指针走到最后，慢指针刚好在中间！！！
*/

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
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
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	res := middleNode(&head)
	fmt.Println(res.Val)
}
