package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 一次性读入，不适合大文件，也不需要打开和关闭文件
	content, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", content) // []byte
	//fmt.Printf("File contents: %s", content)
	fmt.Println(string(content))
}
