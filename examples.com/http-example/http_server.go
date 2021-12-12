package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func login(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./login.html")
		log.Println(t.Execute(w, nil))
	} else if r.Method == "POST" {
		// 表示表单提交数据
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
