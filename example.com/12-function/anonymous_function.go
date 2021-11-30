package main

import (
	"flag"
	"fmt"
)

var (
	Fun1 = func(n1 int, n2 int) int { // 全局匿名函数
		return n1 + n2
	}
)

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	func(s string) {
		fmt.Println(s)
	}("匿名函数")

	// 匿名函数变量
	a := func(s string) {
		fmt.Println(s)
	}

	a("匿名函数变量")

	// 匿名函数做回调函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

	// 匿名函数封装到 map
	var skillParam = flag.String("skill", "", "skill to perform")
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

	// go run main.go --skill=fly
}
