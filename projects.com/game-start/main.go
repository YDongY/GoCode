package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func init() {

}

func main() {
	fmt.Println(runtime.NumCPU())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Golang")
	})
	http.ListenAndServe(":3000", nil)
}
