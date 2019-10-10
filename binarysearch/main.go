package main

import "fmt"

// binarySearch 二分查找
func binarySearch(a []int, val int) int {
	length := len(a)
	left, right := 0, length
	for left < right {
		mid := left + (right-left)/2
		if a[mid] == val {
			return mid
		} else if a[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(a, 8))
}
