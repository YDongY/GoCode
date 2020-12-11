package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 404 错误处理
		return
	}
	switch r.Header.Get("Accept") { // 响应以不同类型的内容
	case "application/json":
		// 设置报头，发送 JSON
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		// 发送 JSON 数据
		w.Write([]byte(`{"name":"jack","age":"20"}`))
	case "text/html":
		w.Header().Set("Content-Type", "text/html;charset=utf-8")
		w.Write([]byte("<h1>hello world</h1>")) // 生成的 HTTP 响应包含状态、报头和响应体
	default:
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Write([]byte("hello world"))
	}

	// cmd : curl -is -H "Accept:application/json" http://localhost:8000/
}

func handler2(w http.ResponseWriter, r *http.Request) {
	// 响应不同类型的请求
	if r.URL.Path != "/" {
		http.NotFound(w, r) // 404 错误处理
		return
	}
	switch r.Method { // 响应以不同类型的内容
	case "GET":
		w.Write([]byte("Get request"))
	case "POST":
		w.Write([]byte("Post request"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	// 获取 https://www.google.com/?q=golang
	for k, v := range r.URL.Query() {
		fmt.Println(k, v)
	}
}

func postBody(w http.ResponseWriter, r *http.Request) {
	// 获取请求体中的数据
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reqBody)
}

func main() {
	/**
	- 路由器默认将没有指定处理程序的请求定向到／。
	- 路由必须完全匹配。例如，对于向／users 发出的请求，将定向到／，因为这里末尾少了斜杆
	- 路由器不关心请求的类型，而只管将与路由匹配的请求传递给相应的处理程序。

	*/

	// http 包默认使用 ServeMux 来处理路由 ， 因此不支持变量，也不支持正则表达式。
	http.HandleFunc("/", handler)     // 注册对 URL 地址映射进行响应的函数
	http.ListenAndServe(":8000", nil) // 启动一个服务器，这个服务器监昕 localhost 和端口 8000
	// http.ListenAndServeTLS() 创建 HTTPS C TLS ）服务器,必须向它传递证书和密钥文件

}
