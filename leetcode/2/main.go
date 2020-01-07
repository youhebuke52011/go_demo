package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := res
	t := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + t
		tmp := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		res.Next = tmp
		res = tmp
		t = sum / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		sum := l1.Val + t
		tmp := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		res.Next = tmp
		res = tmp
		t = sum / 10
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + t
		tmp := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		res.Next = tmp
		res = tmp
		t = sum / 10
		l2 = l2.Next
	}
	if t != 0 {
		res.Next = &ListNode{
			Val:  t,
			Next: nil,
		}
	}
	return head.Next
}

func main() {
	// (2 -> 4 -> 3) + (5 -> 6 -> 4)
	//l1 := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//l2 := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}
	l1 := &ListNode{
		Val: 5,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 5,
		Next: nil,
	}
	fmt.Println(addTwoNumbers(l1, l2))

}
