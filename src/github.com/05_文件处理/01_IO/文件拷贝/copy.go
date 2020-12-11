package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func isExist() (bool, error) {
	// 判断文件是否存在
	_, err := os.Stat("test.txt")
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func copyFile(dstFilePath string, srcFilePath string) (written int64, err error) {
	// 文件拷贝
	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	// 目标文件以写方式打开,文件不存在自动创建
	dsrFile, err := os.OpenFile(dstFilePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	writer := bufio.NewWriter(dsrFile)
	defer dsrFile.Close()
	return io.Copy(writer, reader)
}

func main() {
	fmt.Println(isExist())
	_, err := copyFile("test_copy.txt", "test.txt")
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println(err)
	}
}
