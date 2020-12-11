package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func testReadFile() {
	// func ReadFile(filename string) ([]byte, error)
	fileBytes, err := ioutil.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileBytes)
	fmt.Println(string(fileBytes))
}

func testWriteFile() {
	b := make([]byte, 0)
	// 文件不存在创建文件
	err := ioutil.WriteFile("text2.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 写文件
	s := "hello world"
	err = ioutil.WriteFile("text2.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func testReadDir() {
	// 列出目录内容
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.IsDir(), file.Mode(), file.Name())
	}
}

func testCopyFile() {
	src, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.OpenFile("text_copy.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}

func testRemove() {
	err := os.Remove("text_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	testRemove()
}
