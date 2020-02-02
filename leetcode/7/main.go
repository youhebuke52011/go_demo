package main

import (
	"fmt"
	"math"
	"strconv"
)

func check(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func reverse(x int) int {
	s := strconv.Itoa(x)
	sb := []byte(s)
	l, r := 0, len(sb)-1
	for l < r {
		for l < r && !check(sb[l]) {
			l += 1
		}
		for l < r && !check(sb[r]) {
			r -= 1
		}
		sb[l], sb[r] = sb[r], sb[l]
		l += 1
		r -= 1
	}
	res, _ := strconv.Atoi(string(sb))
	// [−2^31, 2^31 − 1]
	if !(res >= math.MinInt32 && res <= math.MaxInt32) {
		res = 0
	}
	return res
}

func main() {
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(123))
	fmt.Println(reverse(1534236469))
}
