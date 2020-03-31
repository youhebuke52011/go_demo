package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	res := (*m)[0]
	*m = (*m)[:len(*m)-1]
	return res
}

func minK(nums []int, k int) ([]int, error) {
	tmp := MinHeap(nums)
	h := &tmp
	heap.Init(h)
	res := make([]int, k)
	for i:=0;i<k;i++ {
		t := h.Pop().(int)
		res[i] = t
		heap.Remove(h,0)
	}
	return res,nil
}

func main() {
	fmt.Println(minK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}
