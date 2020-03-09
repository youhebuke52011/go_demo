package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
思路：按照 if ab > ba -> a > b, if ab < ba -> a < b, if ab = ba -> a = b 来排序
 */

type T struct {
	tmp []string
}

func (t T) Len() int {
	return len(t.tmp)
}

func (t T) Less(i, j int) bool {
	if t.tmp[i]+t.tmp[j] < t.tmp[j]+t.tmp[i] {
		return true
	}
	return false
}

func (t T) Swap(i, j int) {
	t.tmp[i],t.tmp[j] = t.tmp[j],t.tmp[i]
}

func largestNumber(nums []int) string {
	res := ""
	t := T{
		tmp:[]string{},
	}
	for _, row := range nums {
		t.tmp = append(t.tmp, strconv.Itoa(row))
	}
	sort.Sort(t)
	if t.tmp[len(t.tmp)-1] == "0" {
		return "0"
	}
	for i := len(t.tmp) - 1; i >= 0; i-- {
		res += t.tmp[i]
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{0, 0}))
}
