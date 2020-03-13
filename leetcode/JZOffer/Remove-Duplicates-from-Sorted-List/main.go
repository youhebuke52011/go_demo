package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
删除链表中重复的元素
*/

func RemoveDuplicatesFromSortedList(head *ListNode) {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	RemoveDuplicatesFromSortedList(&head)
	fmt.Println(head)
}
