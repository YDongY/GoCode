package main

import "fmt"

func main() {
	// '/' 相当于整除
	fmt.Println(10 / 4)          // 2
	fmt.Println(10 * 1.0 / 4)    // 2.5
	fmt.Println(float32(10 / 4)) // 2

	fmt.Println(4 / 10)       // 0
	fmt.Println(4 * 1.0 / 10) // 0.4

	// a % b = a - a / b * b
	fmt.Println(10 % 4)   // 2
	fmt.Println(-10 % -3) // -1
	fmt.Println(10 % -3)  // 1 // 10 - 10 / -3 * -3
	fmt.Println(-10 % 3)  // -1

	a, b := 10, 20
	fmt.Println(a, b)
	fmt.Println(a / b)
	fmt.Println(a * 1.0 / b)

	// ++ --
	i := 1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

	// 不能　--i 或者 ++i 或者 j:=i++ i=i++
	// 即：++ , -- 只能单独使用

	// -------------------------
	func2()
	func3()
}

func func2() {
	allDays := 97
	week := allDays / 7
	days := allDays % 7
	fmt.Printf("还有 %d 个星期零 %d 天 \n", week, days)
}

func func3() {
	tmp := 110.65
	n := 5 * 1.0 / 9 * (tmp - 100)

	fmt.Printf("%f \n", n)

}
