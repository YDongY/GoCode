package main

import (
	"bufio"
	"fmt"
	"os"
)

func write() {
	// 创建文件并写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}

	// 关闭
	defer f.Close()

	str := "hello go\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func coverWrite() {
	// 清空源文件，在写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer f.Close()

	str := "hello world\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func appendWrite(){
	// 追加写入
	f, err := os.OpenFile("notes.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer f.Close()

	str := "你好，世界\r\n"

	// 带缓冲写入
	writer := bufio.NewWriter(f)

	for i := 0; i < 1; i++ {
		writer.WriteString(str)
	}

	// 将缓冲 flush 到磁盘
	writer.Flush()
}

func main() {
	// write()
	coverWrite()
	appendWrite()
}
