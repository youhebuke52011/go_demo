package main

import (
	"container/heap"
	"fmt"
)

// O(nk)复杂度
//func maxSlidingWindow(nums []int, k int) []int {
//	res := []int{}
//	for i := k; i <= len(nums); i++ {
//		res = append(res, getMax(nums[i-k:i]))
//	}
//	return res
//}
//
//func getMax(nums []int) int {
//	tmp := nums[0]
//	for _, row := range nums {
//		if tmp < row {
//			tmp = row
//		}
//	}
//	return tmp
//}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	tmp := (*h)[0]
	*h = (*h)[:len(*h)-1]
	//*h = (*h)[1:]
	return tmp
}

func maxSlidingWindow(nums []int, k int) []int {
	a := make([]int, k)
	copy(a,nums[:k] )
	tmp := Heap(a)
	h := &tmp
	heap.Init(h)
	res := []int{}
	for i := k; i < len(nums); i++ {
		//t := heap.Pop(h).(int)
		t := (*h)[0]
		heap.Remove(h, len(*h)-1)
		//t := h.Pop().(int)
		heap.Push(h, nums[i])
		res = append(res, t)
	}
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
