package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct {

}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("My Handler ServeHTTP!!!")
	w.Write([]byte("hello my handler!!!\n"))
}

func MyHandlerServer() {
	server := http.Server{
		Addr:              ":6767",
		Handler:           &MyHandler{},
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	if err := server.ListenAndServe();err != nil {
		panic(err)
	}
}