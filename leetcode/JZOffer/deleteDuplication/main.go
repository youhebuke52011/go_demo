package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
题目：删除链表中重复的结点
pPre:最晚访问不重复的节点；pCur:当前节点；pNext:当前节点下一个节点
*/

func deleteDuplication(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pPre, pCur, pNext := head, head, head.Next
	for pCur != nil {
		if pCur.Next != nil && pCur.Val == pCur.Next.Val {
			pNext = pCur.Next
			for pNext != nil && pNext.Next != nil && pNext.Val == pNext.Next.Val {
				pNext = pNext.Next
			}
			if pCur == head {
				head = pNext
			} else {
				pPre.Next = pNext
			}
			pCur = pNext.Next
			pNext = pNext.Next
		} else {
			pPre = pCur
			pCur = pCur.Next
		}
	}
	return head
}

func main() {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplication(&head)
	for res != nil {
		fmt.Printf("%d ", head.Val)
		res = res.Next
	}
}
