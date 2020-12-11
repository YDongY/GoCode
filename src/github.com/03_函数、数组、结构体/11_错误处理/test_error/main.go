package main

import (
	"errors"
	"fmt"
	"log"
)

func testCreateError() {
	// 自定义错误
	err := errors.New("this a errors")
	if err != nil {
		fmt.Println(err)
	}
}

func testSetErrorFormat() {
	// 设置错误1
	err := fmt.Errorf("%s-%s", "Error", "Set")
	if err != nil {
		fmt.Println(err)
	}
}

func testFuncReturnError() error {
	// 函数返回错误
	return errors.New("return error")
}

func main() {
	err := testFuncReturnError()
	if err != nil {
		log.Fatal(err)
	}
}
