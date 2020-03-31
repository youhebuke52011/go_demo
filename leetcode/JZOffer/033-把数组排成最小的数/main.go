package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
题目描述

输入一个正整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。 输入 例如输入数组

{3，32，321}

输出

则打印出这三个数字能排成的最小数字为

321323

思路：   ab > ba -> a > b  排序
*/

type MyNums []int

func (m MyNums) Len() int {
	return len(m)
}

func (m MyNums) Less(i, j int) bool {
	t1 := strconv.Itoa(m[i]) + strconv.Itoa(m[j])
	t2 := strconv.Itoa(m[j]) + strconv.Itoa(m[i])
	if t1 < t2 {
		return true
	}
	return false
}

func (m MyNums) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func getMinStr(nums []int) string {
	a := MyNums(nums)
	sort.Sort(a)
	res := ""
	for _, row := range a {
		res += strconv.Itoa(row)
	}
	return res
}

func main() {
	fmt.Println(getMinStr([]int{3, 32, 321}))
}
