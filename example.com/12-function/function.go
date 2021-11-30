package main

import "fmt"

func cal(a, b int) int {
	return a + b
}

func cal2(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func reverse(a int, b int) (int, int) {
	return b, a
}

func reverse2(a int, b int) (x int, y int) { // 返回值命名
	x = b
	y = a
	return // 直接返回 x,y
}

func main() {
	fmt.Println(cal(10, 20))
	fmt.Println(cal2(10, 20))
	fmt.Println(reverse(10, 20))
	fmt.Println(reverse2(10, 20))
}
