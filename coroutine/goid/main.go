package main

import "runtime"

func main() {
	//panic("123")
	buf := make([]byte,64)
	stk := buf[:runtime.Stack(buf,false)]
	println(string(stk))
}
