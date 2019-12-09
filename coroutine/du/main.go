package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {

	WalkDirMain()
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func WalkDirMain() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			wg.Add(1)
			go WalkDir(root, fileSizes, &wg)
		}
		//defer close(fileSizes)
	}()

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	// 方式一  range方式接收
	//for fileSize := range fileSizes {
	//	nfiles += 1
	//	nbytes += fileSize
	//}

//  方式二  select 并+定时print 接收
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(5 * time.Second)
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles += 1
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
