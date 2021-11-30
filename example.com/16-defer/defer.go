package main

import "fmt"

func sum(n1 int, n2 int) int {

	// 当执行到 defer 时，暂时不执行，会将 defer 后面的语句压入到栈中
	// 当函数执行完毕，再从栈中依次弹出执行
	defer fmt.Println("n1 =", n1)
	defer fmt.Println("n2 =", n2)
	res := n1 + n2

	fmt.Println("res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println(res)
}

// res= 30
// n2 = 20
// n1 = 10
// 30