package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 || x == 1  {
		return x
	}
	l, r := 1, x
	for l<r {
		mid := l + (r-l)/2
		// 8/4  4, 9/4 4
		if x / mid >= mid {
			l = mid + 1
		} else {
			//
			r = mid
		}
	}
	return r-1
}

func main() {
	fmt.Println(mySqrt(9))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(10))
	fmt.Println(mySqrt(17))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(3))
	//a,b := 1,2
	//a = a+b
	//b = a-b
	//a = a-b
	//fmt.Println(a,b)
}

