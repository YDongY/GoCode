package main

import (
	"errors"
	"fmt"
)

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()

	num1 := 10
	num2 := 0

	// 错误处理 panic ,recover,defer
	fmt.Println(num1 / num2) // panic: runtime error: integer divide by zero
}

func readFile(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func test2() {
	err := readFile("config.ini2")
	if err != nil {
		panic(err) // 程序终止
	}

	fmt.Println("...test2...")
}

func main() {
	test()
	fmt.Println("...end...")

	//　自定义错误 errors.New("错误说明")
	test2()
}
