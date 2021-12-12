package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type ProxyHTTPSServer struct{}

func (p *ProxyHTTPSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s %s\n", r.Method, r.Host, r.RemoteAddr)

	// 连接到源服务器 443 端口的 TCP
	serverConn, err := net.DialTimeout("tcp", r.URL.Host, 10*time.Second)
	if err != nil {
		log.Printf("Dial error:%v\n", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	// 给客户端返回 200 ，表示代理与源服务器连接已建立
	w.WriteHeader(http.StatusOK)

	// 劫持 HTTP 服务器维护的下层的 TCP 连接
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	clientConn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Client <---TCP---> Proxy Server <---TCP---> Server
	defer serverConn.Close()
	defer clientConn.Close()

	go io.Copy(clientConn, serverConn)
	io.Copy(serverConn, clientConn)
}

func main() {
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", &ProxyHTTPSServer{}))
}
