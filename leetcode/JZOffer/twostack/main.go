package main

import "fmt"

/**
题目: 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型
思路: s1 用作压栈,s2用作出栈   出栈时如果s2还有元素则直接返回,反之把s1的元素压栈到s2,并返回s2最后的元素
*/

type queue struct {
	s1, s2 []int
}

func (q *queue) Push(v int) {
	q.s1 = append(q.s1, v)
}

func (q *queue) Pop() int {
	if len(q.s1) == 0 && len(q.s2) == 0 {
		return -1
	}
	n1 := len(q.s1)
	n2 := len(q.s2)
	res := -1
	if n2 != 0 {
		res =  q.s2[n2-1]
		q.s2 = q.s2[:n2-1]
	} else {
		for i := n1 - 1; i >= 0; i-- {
			q.s2 = append(q.s2, q.s1[i])
		}
		q.s1 = []int{}
		res =  q.s2[n1-1]
		q.s2 = q.s2[:n1-1]
	}
	return res
}

func main() {
	q := &queue{
		s1: []int{},
		s2: []int{},
	}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
