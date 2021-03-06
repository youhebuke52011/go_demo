package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	n := len(height)
	if n < 2 {
		return 0
	}
	l, r := 0, n-1
	res := math.MinInt32
	for l < r && l < n && r < n {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		tmp := h * (r - l)
		if tmp > res {
			res = tmp
		}

		if height[l] < height[r] {
			l += 1
		} else {
			r -= 1
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 1}))
	fmt.Println(maxArea([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

	// for k, v := range HPTranMap {
	// 	kb := []byte(k)
	// 	vb := []byte(v)
	// 	if kb[0] == ' ' || kb[len(kb)-1] == ' ' || vb[0] == ' ' || vb[len(vb)-1] == ' ' {
	// 		fmt.Println(k, v)
	// 	}
	// }

}
