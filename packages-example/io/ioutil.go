package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Write() {
	err := ioutil.WriteFile("file.txt", []byte("Hello Golang"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Read() {
	bytes, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}

func main() {
	Write()
	Read()
}
