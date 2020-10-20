package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v \n", file)

	defer file.Close()

	// 创建带缓冲的 Reader
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读取到换行结束
		if err == io.EOF {                  // 表示文件末尾
			break
		}
		fmt.Printf(str)
	}
	fmt.Println("done")
}
