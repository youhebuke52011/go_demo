package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i:=0;i<10;i++{
			ch <- i
			if i%2==0 {
				fmt.Printf("g1:%d\n",i)
			}
		}
		wg.Done()
	}()

	go func() {
		for i:=0;i<10;i++{
			<- ch
			if i%2!=0 {
				fmt.Printf("g2:%d\n",i)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
