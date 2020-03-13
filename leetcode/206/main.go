package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var curHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = curHead
		curHead = head
		head = tmp
	}
	return curHead
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	res := reverseList(&head)
	for res != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
}
