package main

type LinkNode struct {
	Val  int
	Next *LinkNode
}

// 长链表先走，实现右对齐
func firstCommon(h1, h2 *LinkNode) *LinkNode {
	a, b := h1, h2
	for a != nil && b != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next

		if a == nil && h2 != nil {
			a = h2
		}

		if b == nil && h1 != nil {
			b = h1
		}
	}
}

func main() {

}
