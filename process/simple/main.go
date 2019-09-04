package main

import (
	"fmt"
	"sync"
)

func main() {
	//r1, r2, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	//fmt.Println("r1:", r1)
	//fmt.Println("r2:", r2)

	num := 100000000
	//c := make(chan bool)
	//mutex := sync.Mutex{}
	//mutex := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		//mutex.Lock()
		for i := 100000000; i > 0; i-- {
			num++
		}
		//mutex.Unlock()
		//c<-true
		wg.Done()
	}()

	go func() {
		//<-c
		//mutex.Lock()
		for j := 100000000; j > 0; j-- {
			num--
		}
		//mutex.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(num)
}
