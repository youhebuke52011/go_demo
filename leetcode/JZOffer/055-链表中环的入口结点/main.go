package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
lc:142
思路：
1. 快慢指针相遇
2. 相遇后慢指针再跑一圈再相遇就知道这个环多大k
3. 然后就演变成求倒数第N-k个节点，重置快慢指针，快指针先跑k步，然后慢指针出发，再次出发就是环的入口处了
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next.Next
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
