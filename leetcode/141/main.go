package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
判断链表是否有环
 */

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {
	h2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}
	tail := ListNode{
		Val:  -4,
		Next: &h2,
	}
	h2.Next.Next = &tail
	head := ListNode{
		Val:  1,
		Next: &h2,
	}
	fmt.Println(hasCycle(&head))
}
