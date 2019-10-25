package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Handlers struct {
}

func (h Handlers) ResAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("res")
	w.Write([]byte("res\n"))
}

func Do(w http.ResponseWriter, r *http.Request) {
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	fmt.Println(strings.Join(parts, "|"))
	action := ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}
	fmt.Println(action)
	handle := &Handlers{}
	controller := reflect.ValueOf(handle)
	method := controller.MethodByName(action)
	rr := reflect.ValueOf(r)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, rr})
}

func UpgradeServer() {
	http.Handle("/do/", http.HandlerFunc(Do))
	http.HandleFunc("/hello", Hello)
	if err := http.ListenAndServe(":6767", nil); err != nil {
		panic(err)
	}
}
