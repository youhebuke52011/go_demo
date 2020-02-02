package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: head,
	}
	tmpMap := map[int]*ListNode{}
	i := 0
	for node := head; node != nil; node = node.Next {
		tmpMap[i] = node
		i += 1
	}
	newN := len(tmpMap) - n
	if preNode, ok := tmpMap[newN-1]; ok {
		if nextNode, ok := tmpMap[newN+1]; ok {
			// n在链表中间
			preNode.Next = nextNode
		} else {
			// n是最后一个元素
			preNode.Next = nil
		}
		tmpMap[newN].Next = nil
	} else {
		if _, ok := tmpMap[newN+1]; ok {
			res.Next = tmpMap[newN+1]
		} else {
			res.Next = nil
		}
	}
	return res.Next
}

func main() {
	//node := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	head := removeNthFromEnd(node, 3)
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ", head.Val)
	}
}
