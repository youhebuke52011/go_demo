package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// func detectCycle(head *ListNode) *ListNode {
// 	tmpMap := map[*ListNode]bool{}
// 	for head != nil {
// 		if _, ok := tmpMap[head]; ok {
// 			return head
// 		}
// 		tmpMap[head] = true
// 		head = head.Next
// 	}
// 	return nil
// }

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 无环
	if fast == nil || fast.Next == nil {
		return nil
	}

	var k int
	for slow = slow.Next; slow != fast; slow = slow.Next {
		k++
	}

	slow = head
	fast = head.Next
	for slow != fast {
		if k <= 0 {
			slow = slow.Next
		}
		k--
		fast = fast.Next
	}
	return slow
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
	res := detectCycle(&head)
	fmt.Println(res.Val)
}
