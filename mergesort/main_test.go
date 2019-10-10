package main

import (
	"go_demo/common"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	a := common.GenerateSlice(10000000)
	MergeSort(a)
}
