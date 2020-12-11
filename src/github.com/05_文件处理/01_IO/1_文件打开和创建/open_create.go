package main

import (
	"fmt"
	"os"
)

func testCreate() {
	// Create(name string) (*File, error) ：读写，文件存在清空内容，文件不存在创建
	file, err := os.Create("src/github.com/05_文件处理/01_IO/1_文件打开和创建/test.txt")
	if err != nil {
		fmt.Println("error create")
		return
	}
	defer file.Close()
}

func testOpen() {
	// Open(name string) (*File, error) ：只读
	file, err := os.Open("src/github.com/05_文件处理/01_IO/1_文件打开和创建/test.txt")
	if err != nil {
		fmt.Println("error create")
		return
	}
	defer file.Close()
}

func testOpenFile() {
	// OpenFile(name string, flag int, perm FileMode) (*File, error)
	file, err := os.OpenFile("src/github.com/05_文件处理/01_IO/1_文件打开和创建/test.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("error create")
		return
	}
	defer file.Close()
}

func main() {
	testCreate()
}
