package channel

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestMemo_Get(t *testing.T) {
	//http://gopl.io
	urls := []string{
		"https://golang.org", "https://godoc.org", "https://play.golang.org",
		"https://golang.org", "https://godoc.org", "https://play.golang.org",
		"https://golang.org", "https://godoc.org", "https://play.golang.org",
		"https://golang.org", "https://godoc.org", "https://play.golang.org",
	}
	var wg sync.WaitGroup
	m := New(httpGetBody)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
