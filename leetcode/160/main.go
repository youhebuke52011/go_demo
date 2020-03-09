package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
思路：
a1 a2 a3 a4 a5 b1 b2 b3 b4 b5 b6
b1 b2 b3 b4 b5 b6 a1 a2 a3 a4 a5

4,1,8,4,5 5,0,1,8,4,5
5,0,1,8,4,5 4,1,8,4,5
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	linkHeadA, linkHeadB := false, false
	for A != nil && B != nil {
		// 找到了
		if A == B {
			return A
		}
		A, B = A.Next, B.Next

		if A == nil && !linkHeadB {
			A = headB
			linkHeadB = true
		}

		if B == nil && !linkHeadA {
			B = headA
			linkHeadA = true
		}
	}
	return nil
}

func main() {

}
