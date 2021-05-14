package main

import (
	"fmt"
	"unsafe"
)

// `逃逸分析`决定了一个变量是分配在栈上还是分配在堆上
// 逃逸分析  go build-gcflags '-m' main.go
// gdb调试  go build -gcflags "-N -l" main.go
// br 行数，run, c,n,info locals,info goroutines,whatis 变量,
func main() {
	s := "you can you up!"
	// sb := []byte(s)
	sb := Str2bytes(s)

	fmt.Println(sb)
	fmt.Println(s)
	// sb[0] = 'K'
	// fmt.Printf("%c\n",sb[0])

	i := 10
    ip := &i
	fp := (*string)(unsafe.Pointer(ip))
	fmt.Printf("%v\n", i)
    *fp = "11"
	fmt.Println(i)
	

	t := Teacher{"ttt", 20}
    pt := unsafe.Pointer(&t)
    name := (*string)(unsafe.Pointer(pt))
	fmt.Println("name:", *name)
	age := (*int)(unsafe.Pointer(uintptr(pt) + unsafe.Offsetof(t.age)))
	fmt.Println("age:", *age)
}

type Teacher struct {
    name string
    age  int
}

// Str2bytes 高性能string->[]byte
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str 高性能[]byte->string
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
