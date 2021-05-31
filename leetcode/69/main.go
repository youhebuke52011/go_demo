package main

import "fmt"

/**
x的平方跟
*/

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	l, r := 1, x
	for l < r {
		// 不能使用(right+left)，因为有可能整数溢出
		mid := l + (r-l)/2
		// 8/4  4, 9/4 4
		// 找到中间的数: mid = start + (end - start) / 2 而且: mid * mid <= x && (mid + 1) * (mid + 1) > x 但是这么写超出了Integer的范围了:
		// 所以变成下面这样
		if x/mid >= mid {
			l = mid + 1
		} else {
			//
			r = mid
		}
	}
	return r - 1
}

func main() {
	// fmt.Println(mySqrt(9))
	// fmt.Println(mySqrt(8))
	// fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(17))
	// fmt.Println(mySqrt(1))
	// fmt.Println(mySqrt(0))
	// fmt.Println(mySqrt(2))
	// fmt.Println(mySqrt(3))
	//a,b := 1,2
	//a = a+b
	//b = a-b
	//a = a-b
	//fmt.Println(a,b)
}
