package main

import (
	"go_demo/common"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	a := common.GenerateSlice(10000000)
	QuickSort(a)
}
