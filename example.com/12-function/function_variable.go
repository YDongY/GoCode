package main

import "fmt"

func funcType() string {
	return "函数也是一种类型"
}

func getFunc(funcvar func() string) string {
	return funcvar()
}

func returnFunc() func(int) int { // 有参有返回值的函数作为返回值
	return func(i int) int {
		return i
	}
}

func main() {
	f := funcType // 函数作为变量

	fmt.Println(f())

	fmt.Println(getFunc(f)) // 函数作为参数

	fc := returnFunc() // 函数作为返回值
	fmt.Println(fc(10))

	// 等价于
	// fmt.Println(returnFunc()(10))
}