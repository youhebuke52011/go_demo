package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
)

func main() {
	backend := "127.0.0.1:6767"
	target, err := url.Parse(backend)
	if err != nil {
		fmt.Printf("backend %s parse errr : %v\n", backend, err)
	}

	reserveProxy := httputil.NewSingleHostReverseProxy(target)
	reserveProxy.Transport = Transport{cfg}
	c.Request.Host = target.Host
	c.Request.URL.Path = newReqpath

	reserveProxy.ServeHTTP(c.Writer, c.Request)
}
