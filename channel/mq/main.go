package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i:=0;i<10;i++{
			ch <- i
			fmt.Printf("send:%d\n", i)
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<10;i++{
			tmp := <-ch
			fmt.Printf("recv:%d\n", tmp)
		}
		wg.Done()
	}()

	wg.Wait()
}
