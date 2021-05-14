package main

import "fmt"

/**
指数幂的所有边界包括:
- 指数为0的情况，不管底数是多少都应该是1
- 指数为负数的情况，求出的应该是其倒数幂的倒数
- 指数为负数的情况下，底数不能为0
*/

func Pow(base float64, exp int) (float64, error) {
	if exp == 0 {
		return 1, nil
	}
	if base == 0 && exp < 0 {
		return -1, nil
	}
	var res float64
	if exp > 0 {
		res = PowNormal(base, exp)
	} else {
		res = PowNormal(base, -exp)
		res = 1 / res
	}
	return res, nil
}

func PowNormal(base float64, exp int) float64 {
	res, tmp := 1.0, base
	for exp != 0 {
		// 奇数
		fmt.Println(exp, tmp, exp&1, exp>>1)
		if exp&1 == 1 {
			res *= tmp
		}
		// 2->4->16->256
		tmp *= tmp
		exp >>= 1
	}
	return res
}

func main() {
	res, _ := Pow(2, 5)
	fmt.Println(res)
}
