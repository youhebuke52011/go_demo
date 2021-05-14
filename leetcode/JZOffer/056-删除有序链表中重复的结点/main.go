package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
例如，链表1->2->3->3->4->4->5 处理后为 1->2->5
题目：删除链表中重复的结点
pPre:最晚访问不重复的节点；pCur:当前节点；pNext:当前节点下一个节点
*/

// func deleteDuplication(head *ListNode) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	pPre, pCur, pNext := head, head, head.Next
// 	for pCur != nil {
// 		if pCur.Next != nil && pCur.Val == pCur.Next.Val {
// 			pNext = pCur.Next
// 			for pNext != nil && pNext.Next != nil && pNext.Val == pNext.Next.Val {
// 				pNext = pNext.Next
// 			}
// 			if pCur == head {
// 				head = pNext
// 			} else {
// 				pPre.Next = pNext
// 			}
// 			pCur = pNext.Next
// 			pNext = pNext.Next
// 		} else {
// 			pPre = pCur
// 			pCur = pCur.Next
// 		}
// 	}
// 	return head
// }

func deleteDuplication(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	pPre, pCur, pNext := head, head, head.Next
	for pNext != nil {

		if pCur == head && pCur.Val == pNext.Val {
			// 头节点就有重复元素
			for pNext != nil && pCur.Val == pNext.Val {
				pCur = pCur.Next
				pNext = pNext.Next
			}
			if pNext == nil {
				return nil
			}
			pCur = pCur.Next
			pNext = pNext.Next
			head = pCur

		} else if pCur.Val == pNext.Val {
			// 中间或结尾节点有重复元素
			for pPre.Next != pCur {
				pPre = pPre.Next
			}
			for pNext != nil && pCur.Val == pNext.Val {
				pNext = pNext.Next
			}
			pCur = pPre
			pCur.Next = pNext
		} else {
			pCur = pCur.Next
			pNext = pNext.Next
		}
		fmt.Printf("%d %d ", pPre.Val, pCur.Val)
		if pNext != nil {
			fmt.Println(pNext.Val)
		} else {
			fmt.Println()
		}
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
					Val: 3,
					Next: &ListNode{
						Val:  4,
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
	fmt.Println()
}
