// 展示内存增长和pprof，并不是泄露
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 运行一段时间：fatal error: runtime: out of memory
//func main() {
//	//开启pprof
//	go func() {
//		ip := "0.0.0.0:6060"
//		if err := http.ListenAndServe(ip, nil); err != nil {
//			fmt.Printf("start pprof failed on %s\n", ip)
//			os.Exit(1)
//		}
//	}()
//
//	tick := time.Tick(time.Second / 100)
//	var buf []byte
//	for range tick {
//		buf = append(buf, make([]byte, 1024*1024)...)
//		time.Sleep(10 * time.Millisecond)
//	}
//}

const (
	port = ":12600"
)

type server struct {
}

func (s *server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		rw.Write([]byte("Hello Server!!!"))
	} else if req.URL.Path == "/for" {
		s.forPProf(rw, req)
	}
	return
}

func (s *server) forPProf(rw http.ResponseWriter, req *http.Request)  {
	ticker := time.NewTicker(time.Second * time.Duration(5))
	var buf []byte
	buf = append(buf, make([]byte, 1024*1024)...)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Doing ticker!!!")
				//return
			case <-req.Context().Done():
				fmt.Println("Context Done!!!")
				return
				//default:
				//	rw.Write([]byte("test for pprof!!!"))
			}
		}
	}()
	rw.Write([]byte("test for pprof!!!"))
	return
}

func main() {
	s := &server{}
	err := http.ListenAndServe(port, s)
	if err != nil {
		return
	}
}
