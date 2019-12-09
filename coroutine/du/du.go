package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func WalkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			WalkDir(filepath.Join(dir, entry.Name()), fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	//acquire token
	sema <- struct{}{}
	//release token
	defer func() {
		<-sema
	}()
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return infos
}
