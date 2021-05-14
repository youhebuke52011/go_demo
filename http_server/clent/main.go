package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:6767", nil)
	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	i := 0
	for i < 1000 {
		go func() {
			resp, err := client.Do(req)
			fmt.Println(runtime.NumGoroutine())
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(body))
		}()
		i++
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	time.Sleep(time.Minute)
}
