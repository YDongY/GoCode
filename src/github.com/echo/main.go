package main

import (
	"fmt"
	"os"
	"time"
)

func func1() {
	var s, sep string // 默认初始化空字符串
	/*
		$ go run echo.go a b c
		[/tmp/go-build606173642/b001/exe/echo a b c]
	*/

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func func2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func func3() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}

func main() {
	func1()
	func2()
	func3()
}
