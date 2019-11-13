package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type MyHandler struct{}

type RequestPayload struct {
	ProxyCondition string `json:"proxy_condition"`
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Fatalln("My Handler ServeHTTP!!!")
	if err := Redirect(w, r); err != nil {
		log.Fatalf("ServeHTTP Redirect Error:%v", err)
		panic(err)
	}
}

//Redirect 反向代理 流量转发
func Redirect(w http.ResponseWriter, r *http.Request) error {
	// 解析header proxy_condition
	requestPayload, err := parseRequestBody(r)
	if err != nil {
		return err
	}
	// 根据proxy_condition获取endpoint
	targetUrl := getProxyUrl(requestPayload.ProxyCondition)
	// 流量转发
	serveReverseProxy(targetUrl, w, r)
	return nil
}

// parseRequestBody 解析request header proxy_condition
func parseRequestBody(r *http.Request) (*RequestPayload, error) {
	// io.ReadCloser类型，意味着一旦全部读取完成，就无法进行第二次读取
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	// 把已经读取出来的body数据，使用NopCloser重新包装成io.ReadCloser
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	decoder := json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))

	var requestPayload *RequestPayload
	err = decoder.Decode(requestPayload)
	if err != nil {
		panic(err)
	}

	return requestPayload, nil
}

// 根据配置获取endpoints
func getProxyUrl(proxyCondition string) string {
	//endPoints := strings.Split(os.Getenv("end_points"), ",")
	endPointsMap := map[string]string{
		"ep1": "127.0.0.1:6768",
		"ep2": "127.0.0.1:6769",
		"ep3": "127.0.0.1:6770",
	}
	defaultEndPoint := endPointsMap["ep1"]

	for name, url := range endPointsMap {
		if name == proxyCondition {
			return url
		}
	}
	return defaultEndPoint
}

// serveReverseProxy 反向代理
func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	targetUrl, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)

	r.URL.Host = targetUrl.Host
	r.URL.Scheme = targetUrl.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = targetUrl.Host
	w.Header().Set("Cache-Control", "max-age=30, no-cache")

	proxy.ServeHTTP(w, r)
}

func MyHandlerServer() {
	server := http.Server{
		Addr:              ":6767",
		Handler:           &MyHandler{},
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
