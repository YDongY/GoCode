package main

import "fmt"

func assignmentOperator() {
	// ---------------- 赋值运算符 ------------------------
	// = += -= *= /=
	fmt.Println(".......")
	a := 10
	b := 20

	c := a
	a = b
	b = c
	fmt.Printf("%d,%d \n", a, b)

	n, m := 10, 20

	x := n + m
	n = x - n
	m = x - n

	fmt.Printf("%d,%d \n", n, m)
}

func arithmeticOperator() {
	// ----------------- 算数运算符 -----------------
	// '/' 相当于整除
	fmt.Println(10 / 4)          // 2
	fmt.Println(10 * 1.0 / 4)    // 2.5
	fmt.Println(float32(10 / 4)) // 2
	fmt.Println(4 / 10)          // 0
	fmt.Println(4 * 1.0 / 10)    // 0.4

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
}

func logicalOperators() {
	// ----------------- 逻辑运算符 -----------------
	fmt.Println(1 == 1 && 2 > 1) // true
	fmt.Println(1 == 1 && 2 > 3) // false
	fmt.Println(1 == 1 || 2 > 3) // true
	fmt.Println(1 != 1 || 2 > 3) // false

	fmt.Println(!true)     // false
	fmt.Println(1 != 1)    // false
	fmt.Println(!(1 != 1)) // true

	// 如果 && 左边为假，则不判断右边
	// 如果 || 右边为真，则不判断右边
}

func relationalOperator() {
	// ----------------- 关系运算符 ---------------------
	fmt.Println(1 == 1)  // true
	fmt.Println(1 != -1) // true
	fmt.Println(1 < -1)  // false
	fmt.Println(1 > -1)  // true
	fmt.Println(1 >= -1) // true
	fmt.Println(1 <= -1) // false

	flag := 1 > 1
	fmt.Println(flag) // false

	flag2 := 1 == 1
	fmt.Println(flag2) // true
}

func bitOperator()  {
	// ---------------- 位运算符 -------------------
	var a int8 = 3
	var b int8 = -2

	// 左移
	fmt.Println(a << 1) // 6
	fmt.Println(b << 1) // -4

	// 右移
	fmt.Println(a >> 1) // 1
	// 负数右移最小值为 -1，操作原码
	fmt.Println(b >> 1) // -1

	// 负数按照补码计算: | & ^

	fmt.Println(a | b) // -1

	fmt.Println(a & b) // 2

	fmt.Println(a ^ b) // -3

	fmt.Println(a &^ b) // 1
}

func main() {

}
