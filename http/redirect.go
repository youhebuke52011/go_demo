package main

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RoundTripper interface {
	RoundTrip(*http.Request) (*http.Response, error)
}

type LogRedirects struct {
	Transport http.RoundTripper
	Urls      []string
	proxy     *url.URL
	header    string
}

type Url struct {
	Url           string
	Status, Delta int
}

type adtrackschema struct {
	Url Url
}



func (l *LogRedirects) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t := l.Transport
	if t == nil {
		t = http.DefaultTransport
	}
	resp, err = t.RoundTrip(req)
	if err != nil {
		return
	}
	switch resp.StatusCode {
	case http.StatusMovedPermanently, http.StatusFound, http.StatusSeeOther, http.StatusTemporaryRedirect:
		//log.Println("Request for", req.URL, "redirected with status", resp.StatusCode)
		l.Urls = append(l.Urls, req.URL.String()+" "+resp.Status)
	default:
		l.Urls = append(l.Urls, req.URL.String()+" "+resp.Status)
	}
	return
}

// 记录各个跳转详情
func (l *LogRedirects) RoundTripDetail(req *http.Request) (resp *http.Response, err error) {
	proxy := l.proxy
	t := http.RoundTripper(&http.Transport{Proxy: http.ProxyURL(proxy)})
	now := time.Now()
	// 不传Referer
	req.Header.Del("Referer")
	resp, err = t.RoundTrip(req)
	if err != nil {
		// 落地页直接写入
		if isLandingPage(req.URL) {
			l.Urls = append(l.Urls, &adtrackschema.Url{
				Url: req.URL.String(),
				Status: 200,
				Delta:  0,
			})
		}
		return
	}
	// 记录耗时
	delta := int64(time.Since(now) / time.Millisecond)
	// 记录最后一条的Header
	l.header = ""
	for k, v := range resp.Header {
		l.header += k + ":" + strings.Join(v, ",") + ";"
	}
	// 添加层级
	l.Urls = append(l.Urls, &adtrackschema.Url{Url: req.URL.String(),
		Status: int64(resp.StatusCode),
		Delta:  delta})
	return
}

func isLandingPage(u *url.URL) bool {
	return true
}

func main() {

}
