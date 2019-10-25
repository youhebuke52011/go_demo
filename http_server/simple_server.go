package main

import "net/http"

/**
最简单的http server
 */

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello http server!\n"))
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home!\n"))
}

func SimpleServer() {
	http.Handle("/home", http.HandlerFunc(Home))
	http.HandleFunc("/hello", Hello)
	if err := http.ListenAndServe(":6767", nil); err != nil {
		panic(err)
	}
}
