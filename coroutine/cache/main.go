package main

import (
	"fmt"
	"go_demo/coroutine/cache/memo"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	urls := []string{
		"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io",
		"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io",
	}
	m := memo.New(httpGetBody)
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}
