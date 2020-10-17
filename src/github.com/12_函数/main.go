package main

import (
	"fmt"
)

// func 函数名 (参数列表) (返回值列表){
// 	执行语句...
// 	return
// }

// func cal(a int, b int) int {
// 	return a + b
// }

// 简写
func cal(a, b int) int {
	return a + b
}

func reverse(a int, b int) (int, int) {
	return b, a
}

func reverse2(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func reverse3(a int, b int) (x int, y int) { // 返回值命名
	x = b
	y = a
	return // 直接返回 x,y
}

// 传递地址
func modify(n *int) {
	fmt.Println(n)  // 0xc0000160e0
	fmt.Println(&n) // 0xc00000e030
	fmt.Println(*n) // 100
	*n += 1
	fmt.Println(*n) // 101
}

func funcType() string {
	return "函数也是一种类型"
}

func getFunc(funcvar func() string) string {
	return funcvar()
}

// func getSum(n1 int, n2 int) int {
// 	return n1 + n2
// }

// func myFun(funvar func(int, int) int, num1 int, num2 int) int {
// 	return funvar(num1, num2)
// }

func sumArgs(n1 int, args ...int) int { // 函数接收可变参数
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

func main() {
	fmt.Println(cal(10, 20))
	fmt.Println(reverse(10, 20))
	sum, _ := reverse2(10, 20) // 下划线占位符
	fmt.Println(sum)
	fmt.Println(reverse3(10, 20))

	num := 100
	modify(&num) // 传递地址，num 值会发生改变
	fmt.Println(num)

	f := funcType // 函数作为变量

	fmt.Println(f())

	fmt.Println(getFunc(f)) // 函数作为参数

	// res := myFun(getSum, 10, 20) // 函数作为参数
	// fmt.Println(res)

	// ----------------- 自定义类型 --------------------------
	type myInt int

	var i myInt
	i = 10
	fmt.Printf("%T,%d \n", i, i) // main.myInt,10

	type funcDemo func(int, int) int

	// 求和
	rest := sumArgs(1, 2, 3, 4, 5, 6) // 可变参数
	fmt.Println(rest)
}
