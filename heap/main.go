package main

import "fmt"

type MaxHeap struct {
	Heap     []int
	Count    int
	Capacity int
}

// 初始化heap,空堆
func InitHeap(size int) MaxHeap {
	maxHeap := MaxHeap{
		// 0不存数据
		Heap:     make([]int, size+1),
		Count:    0,
		Capacity: size + 1,
	}
	return maxHeap
}

// 构造函数, 通过一个给定数组创建一个最大堆
// 该构造堆的过程, 时间复杂度为O(n)
func InitHeapByHeapify(items []int) MaxHeap {
	size := len(items)
	maxHeap := InitHeap(size)

	for i, row := range items {
		maxHeap.Heap[i+1] = row
		maxHeap.Count++
	}

	// Heapify过程,从第一个有叶子节点的节点开始(maxHeap.Count / 2)
	for i := maxHeap.Count / 2; i > 0; i-- {
		maxHeap.ShiftDown(i)
	}
	return maxHeap
}

// 获取堆大小
func (h *MaxHeap) Size() int {
	return h.Count
}

// 判断堆是否为空
func (h *MaxHeap) IsEmpty() bool {
	if h.Count == 0 {
		return true
	}
	return false
}

// ShiftUp (假设是个最大堆)向插入一个value（位置为k）,并和父节点f（k/2）比较,if f<value ; swap(k,2/k); 一直到顶; 维持最大堆
// 最大（小）堆插入维护
func (h *MaxHeap) ShiftUp(k int) {
	for ; k > 1 && h.Heap[k] > h.Heap[k/2]; k /= 2 {
		h.Heap[k], h.Heap[k/2] = h.Heap[k/2], h.Heap[k]
	}
}

// ShiftDown 最大（小）堆删除维护
func (h *MaxHeap) ShiftDown(k int) {
	//h.Count++
	//h.Heap[1], h.Heap[h.Count] = h.Heap[h.Count], h.Heap[1]
	for 2*k <= h.Count {
		son := 2 * k // swap(son,k)
		// 如果有右孩子,比较下左右孩子哪个大,哪个大就和哪个交换
		if son+1 <= h.Count && h.Heap[son] < h.Heap[son+1] {
			son++
		}
		// 已经符合最大（小）堆要求
		if h.Heap[k] > h.Heap[son] {
			break
		}
		// 交换
		h.Heap[k], h.Heap[son] = h.Heap[son], h.Heap[k]
		k = son
	}
}

// Insert 插入到最后一位
func (h *MaxHeap) Insert(item int) {
	if h.Count+1 > h.Capacity {
		return
	}
	h.Heap[h.Count+1] = item
	h.ShiftUp(h.Count + 1)
	h.Count++
}

// GetMax 获取堆顶元素,即最大值
func (h *MaxHeap) GetMax() int {
	if h.Count <= 0 {
		return 0
	}
	return h.Heap[1]
}

// ExtractMax 取出堆顶元素
func (h *MaxHeap) ExtractMax() int {
	if h.Count <= 0 {
		return 0
	}
	ret := h.Heap[1]
	// 最大值替换到数组末尾的位置
	h.Heap[1], h.Heap[h.Count] = h.Heap[h.Count], h.Heap[1]
	// 数量--
	h.Count--
	// 重新维持最大堆
	h.ShiftDown(1)
	return ret
}

func main() {
	arr := []int{8, 8, 3, 3, 2, 1, 2, 3, 6, 1, 1}
	n := len(arr)

	// 对整个arr数组使用HeapSort1排序
	// HeapSort1, 将所有的元素依次添加到堆中, 在将所有元素从堆中依次取出来, 即完成了排序
	// 无论是创建堆的过程, 还是从堆中依次取出元素的过程, 时间复杂度均为O(nlogn)
	// 整个堆排序的整体时间复杂度为O(nlogn)
	maxHeap := InitHeap(n)
	for _,row := range arr {
		maxHeap.Insert(row)
	}
	for i := 1; i < n; i++ {
		fmt.Print(maxHeap.ExtractMax())
		fmt.Print(" ")
	}
	fmt.Println()


	// 对整个arr数组使用HeapSort2排序
	// HeapSort2, 借助我们的heapify过程创建堆
	// 此时, 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 实践复杂度为O(nlogn)
	// 堆排序的总体时间复杂度依然是O(nlogn), 但是比HeapSort1性能更优, 因为创建堆的性能更优
	maxHeap = InitHeapByHeapify(arr)
	for i := 1; i < n; i++ {
		fmt.Print(maxHeap.ExtractMax())
		fmt.Print(" ")
	}
	fmt.Println()
}
