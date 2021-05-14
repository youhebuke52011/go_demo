package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		prev *ListNode
		slow = head
		fast = head
	)
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	l1 := sortList(head)
	l2 := sortList(slow)
	return merge(l1, l2)
}

func merge(l, r *ListNode) *ListNode {
	var (
		head = &ListNode{}
		cur  = head
	)

	for l != nil && r != nil {
		if l.Val <= r.Val {
			tmp := l.Next
			cur.Next = l
			l.Next = nil
			l = tmp
		} else {
			tmp := r.Next
			cur.Next = r
			r.Next = nil
			r = tmp
		}
		cur = cur.Next
	}

	if l != nil {
		cur.Next = l
	}
	if r != nil {
		cur.Next = r
	}

	return head.Next
}

func main() {
	// node := ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: nil}}}}
	node := ListNode{Val: -1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: nil}}}}}
	res := sortList(&node)
	for res.Next != nil {
		fmt.Printf("%d ", res.Val)
		res = res.Next
	}
	fmt.Println(res.Val)
}
