package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode 21

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	res := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		res = res.Next
	}
	if l1!=nil{
		res.Next = l1
	}
	if l2!=nil{
		res.Next = l2
	}
	return head.Next
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  5,
		Next: nil,
	}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Printf("%d,", res.Val)
		res = res.Next
	}
}
