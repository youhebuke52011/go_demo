package main

import (
	"strings"
	"testing"
)

var s = strings.Repeat("a", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}

func test2() {
	b := Str2bytes(s)
	_ = Bytes2str(b)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
