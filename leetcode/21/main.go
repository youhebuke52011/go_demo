package main

import "fmt"

//递归法 4ms
//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1==nil{return l2}
//	if l2==nil{ return l1}
//	if l1.Val<l2.Val{
//		l1.Next=mergeTwoLists(l1.Next,l2)
//		return l1
//	}else{
//		l2.Next=mergeTwoLists(l1,l2.Next)
//		return l2
//	}
//}
////迭代法 0ms
//func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
//	//设置个哨兵
//	preHead:=new(ListNode)
//	//设置当前位置的滑动指针
//	var current = preHead
//	for l1!=nil && l2!=nil{
//		if l1.Val<=l2.Val{
//			current.Next=l1
//			l1=l1.Next
//		}else {
//			current.Next=l2
//			l2=l2.Next
//		}
//		current = current.Next
//	}
//	if l1==nil{
//		current.Next=l2
//	} else {
//		current.Next=l1
//	}
//	return preHead.Next
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return head.Next
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  5,
		Next: nil,
	}
	res := mergeTwoLists(l1, l2)
	for res != nil {
		fmt.Printf("%d,", res.Val)
		res = res.Next
	}
}
