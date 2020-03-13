package main

import "fmt"

// 逆序对 (3,1) (7,6) ...

func InversionCount(arr []int) int {
	count := 0
	n := len(arr)
	l := 0     // arr[0...n/2)
	r := n / 2 // arr[n/2...n)

	for l < n/2 && r < n {
		if arr[l] > arr[r] {
			count++
			r++
		} else {
			l++
		}
	}

	for l < n/2 {
		l++
	}

	for r < n {
		r++
	}
	return count
}

func main() {
	arr := []int{3, 7, 2, 9, 6, 1, 8}
	fmt.Println(InversionCount(arr))
}
