package main

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HTTPClientGet() {
	resp, err := http.Get("https://www.baidu.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if body, err := io.ReadAll(resp.Body); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", body)
	}
}

func HTTPClientGetWithArgs() {
	apiUrl := "http://127.0.0.1:8080/get"

	data := url.Values{} // URL param
	data.Set("name", "mark")
	data.Set("age", "24")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println(u.String())    // http://127.0.0.1:8080/get?age=24&name=mark

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func HTTPClientPost() {
	apiUrl := "http://127.0.0.1:8080/post"
	// -------- form data --------
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=mark&age=18"
	// -------- json data --------
	contentType := "application/json"
	data := `{"name":"mark","age":18}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func CustomTransport() {
	// 要管理代理、TLS 配置、keep-alive、压缩和其他设置，创建一个 Transport
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://www.baidu.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if body, err := io.ReadAll(resp.Body); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", body)
	}

}

func DefaultHTTPServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func CustomHTTPServer() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        &myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func HTTPServer() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		data := r.URL.Query()
		fmt.Println(data.Get("name"))
		fmt.Println(data.Get("age"))
		answer := `{"status": "ok"}`
		w.Write([]byte(answer))
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
		if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
			// 1. 请求类型是 application/x-www-form-urlencoded 时解析 form 数据
			r.ParseForm()
			fmt.Println("x-www-form-urlencoded:", r.PostForm, r.PostForm.Get("name"), r.PostForm.Get("age")) // 打印 form 数据
		} else if r.Header.Get("Content-Type") == "application/json" {
			// 2. 请求类型是 application/json 时从 r.Body 读取数据
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("application/json:", string(body))
		}
		answer := `{"status": "ok"}`
		w.Write([]byte(answer))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	CustomTransport()
}
