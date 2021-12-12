package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func NewWriter() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 获取 Writer 对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Hello Golang\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func NewReader() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 获取 Reader 对象
	reader := bufio.NewReader(file)
	for {
		line, _, err2 := reader.ReadLine()
		if err2 != nil {
			return
		}
		fmt.Println(string(line))
	}
}

func main() {
	NewReader()
}
