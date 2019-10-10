package main

import "go_demo/common"

/**
把一个数组n分为[l:mid]、[mid:n] ,递归分,直到数组长度==1,然后left、right数组进行merge
merge的过程是两个数组进行比较，借助个长度为len(left)+len(right)的数组result来把两数组比较的更(大)小值放进去
然后递归逆序回去就把各层级的两数组合成一个有序的数组
*/
// MergeSort 归并排序
func MergeSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}
	l, r := 0, length
	//mid := int(length / 2)
	mid := l + (r-l)/2
	//left := make([]int, mid)
	//right := make([]int, length-mid)
	//for i, row := range a {
	//	if i < mid {
	//		left[i] = row
	//	} else {
	//		right[i-mid] = row
	//	}
	//}

	//rl := mergeSort(left)
	//rr := mergeSort(right)
	//result := merge(rl, rr)
	//return merge(mergeSort(left), mergeSort(right))
	return merge(MergeSort(a[l:mid]), MergeSort(a[mid:r]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// left 还有元素
	for ; i < len(left); i++ {
		result[k] = left[i]
		k++
	}

	// right 还有元素
	for ; j < len(right); j++ {
		result[k] = right[j]
		k++
	}
	return result
}

func main() {
	a := common.GenerateSlice(10000000)
	//printSlice(a, "before:")
	//result := MergeSort(a)
	MergeSort(a)
	//printSlice(result, "after:")
}
