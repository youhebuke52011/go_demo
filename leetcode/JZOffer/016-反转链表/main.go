package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 头插法 leetcode:206

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var cur *ListNode
	for head != nil {
		tmp := head.Next
		// 把上个弄好的赋值到head的next
		head.Next = cur
		cur = head
		head = tmp
	}
	return head
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
	res := reverse(&head)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
