package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 长链表先走，实现右对齐  leetcode 160
func getIntersectionNode1(h1, h2 *ListNode) *ListNode {
	fmt.Println(h1.Val)
	a, b := h1, h2
	flaga, flagb := false, false
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a, b = a.Next, b.Next

		if a == nil && !flagb {
			a.Next = h2
			flagb = true
		}

		if b == nil && !flaga {
			b.Next = h1
			flaga = true
		}
	}
	return nil
}

func main() {
	l := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}
	l1 := ListNode{Val: 4, Next: &ListNode{Val: 1, Next: l}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: l}}}
	// fmt.Println(l, l1, l2)
	res := getIntersectionNode1(&l1, &l2)
	fmt.Println(res.Val)
}
