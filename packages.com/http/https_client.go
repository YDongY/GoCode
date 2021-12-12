package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	proxyURL, _ := url.Parse("http://127.0.0.1:8080")
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: tr,
	}

	response, err := client.Get("https://go.dev/")
	if err != nil {
		log.Fatalf("Get error:%v\n", err)
	}
	io.Copy(os.Stdout, response.Body)
}