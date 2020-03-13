package main

import (
	"container/heap"
	"fmt"
)

type lowHeap []int

func (l lowHeap) Len() int {
	return len(l)
}

func (l lowHeap) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l lowHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *lowHeap) Push(x interface{}) {
	*l = append(*l, x.(int))
}

func (l *lowHeap) Pop() interface{} {
	res := (*l)[0]
	*l = (*l)[:len(*l)-1]
	return res
}

func smallestK(arr []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	tmp := lowHeap(arr)
	h := &tmp
	heap.Init(h)
	res := make([]int, k)
	res[0] = h.Pop().(int)
	for i := 1; i < k; i++ {
		res[i] = heap.Remove(h, 0).(int)
	}
	return res

}

func main() {
	//fmt.Println(smallestK([]int{1, 2, 3}, 0))
	//fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
	res := smallestK([]int{62577, -220, -8737, -22, -6, 59956, 5363, -16699, 0, -10603, 64, -24528, -4818, 96, 5747, 2638, -223, 37663, -390, 35778, -4977, -3834, -56074, 7, -76, 601, -1712, -48874, 31, 3, -9417, -33152, 775, 9396, 60947, -1919, 683, -37092, -524, -8, 1458, 80, -8, 1, 7, -355, 9, 397, -30, -21019, -565, 8762, -4, 531, -211, -23702, 3, 3399, -67, 64542, 39546, 52500, -6263, 4, -16, -1, 861, 5134, 8, 63701, 40202, 43349, -4283, -3, -22721, -6, 42754, -726, 118, 51, 539, 790, -9972, 41752, 0, 31, -23957, -714, -446, 4, -61087, 84, -140, 6, 53, -48496, 9, -15357, 402, 5541, 4, 53936, 6, 3, 37591, 7, 30, -7197, -26607, 202, 140, -4, -7410, 2031, -715, 4, -60981, 365, -23620, -41, 4, -2482, -59, 5, -911, 52, 50068, 38, 61, 664, 0, -868, 8681, -8, 8, 29, 412}, 131)
	for _, row := range res {
		fmt.Printf("%d,", row)
	}
}
