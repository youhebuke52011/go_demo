package main

import (
	"fmt"
	"sort"
)

type Vals [][]int

func (v Vals) Len() int {
	return len(v)
}

func (v Vals) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v Vals) Less(i, j int) bool {
	return v[i][0] < v[j][0]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Sort(Vals(intervals))
	n := len(intervals)
	for i := 0; i < n-1; {
		cur, next := intervals[i], intervals[i+1]
		if next[0] <= cur[1] {
			if next[1] > cur[1] {
				cur[1] = next[1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			n--
			continue
		}
		i++
	}
	return intervals
}

func main() {
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
