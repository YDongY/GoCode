package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type ProxyHTTPServer struct {
	Proto string
	Addr  string
	Port  string
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func setProxyHeader(key string, header string, headers http.Header) {
	switch key {
	case "Via":
		if viaSlice, ok := headers[key]; ok {
			header = strings.Join(viaSlice, ", ") + ", " + header
		}
	case "X-Forwarded-For":
		if xffSlice, ok := headers["X-Forwarded-For"]; ok {
			header = strings.Join(xffSlice, ", ") + ", " + header
		}
	}
	headers.Set(key, header)
}

func setViaHeader(viaHeader string, headers http.Header) {
	// 设置 Via
	setProxyHeader("Via", viaHeader, headers)
}

func setXFFHeader(xffHeader string, headers http.Header) {
	// 设置 XFF
	setProxyHeader("X-Forwarded-For", xffHeader, headers)
}

func (p *ProxyHTTPServer) SetRequestHeader(req *http.Request) {
	viaHeader := fmt.Sprintf("%s %s", p.Proto, net.JoinHostPort(p.Addr, p.Port))
	setViaHeader(viaHeader, req.Header)

	host, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		log.Fatalf("SplitHostPort error:%v\n", err)
	}
	setXFFHeader(host, req.Header)

}

func (p *ProxyHTTPServer) SetResponseHeader(w http.ResponseWriter, resp *http.Response) {
	viaHeader := fmt.Sprintf("%s %s", p.Proto, net.JoinHostPort(p.Addr, p.Port))
	setViaHeader(viaHeader, resp.Header)

	setXFFHeader(resp.Request.Host, resp.Header)

	copyHeader(w.Header(), resp.Header)
}

func (p *ProxyHTTPServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("Received request: %s %s %s\n", req.Method, req.RequestURI, req.RemoteAddr)

	// TODO: http / https

	// 1. 设置 HTTP 代理请求头信息
	p.SetRequestHeader(req)

	// 2. 转发给源服务器
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
	}
	defer resp.Body.Close()

	// 3. 设置 HTTP 代理响应头信息
	p.SetResponseHeader(w, resp)

	// 4. 响应源服务器的数据给客户端
	io.Copy(w, resp.Body)
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "Input Your IP Address")
	port := flag.String("port", "8080", "Input Your Port")

	flag.Parse()

	p := &ProxyHTTPServer{
		Proto: "HTTP/1.1",
		Addr:  *ip,
		Port:  *port,
	}
	server := &http.Server{
		Addr:           net.JoinHostPort(p.Addr, p.Port),
		Handler:        p,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
