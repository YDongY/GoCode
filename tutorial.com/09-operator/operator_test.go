package main

import (
	"testing"
)

func TestAssignmentOperator(t *testing.T) {
	// ---------------- 赋值运算符 ------------------------
	// = += -= *= /=
	t.Log(".......")
	a := 10
	b := 20

	c := a
	a = b
	b = c
	t.Logf("%d,%d \n", a, b)

	n, m := 10, 20

	x := n + m
	n = x - n
	m = x - n

	t.Logf("%d,%d \n", n, m)
}

func TestArithmeticOperator(t *testing.T) {
	// ----------------- 算数运算符 -----------------
	// '/' 相当于整除
	t.Log(10 / 4)          // 2
	t.Log(10 * 1.0 / 4)    // 2.5
	t.Log(float32(10 / 4)) // 2
	t.Log(4 / 10)          // 0
	t.Log(4 * 1.0 / 10)    // 0.4

	// a % b = a - a / b * b
	t.Log(10 % 4)   // 2
	t.Log(-10 % -3) // -1
	t.Log(10 % -3)  // 1 // 10 - 10 / -3 * -3
	t.Log(-10 % 3)  // -1

	a, b := 10, 20
	t.Log(a, b)
	t.Log(a / b)
	t.Log(a * 1.0 / b)

	// ++ --
	i := 1
	i++
	t.Log(i)
	i--
	t.Log(i)

	// 不能　--i 或者 ++i 或者 j:=i++ i=i++
	// 即：++ , -- 只能单独使用
}

func TestLogicalOperators(t *testing.T) {
	// ----------------- 逻辑运算符 -----------------
	t.Log(1 == 1 && 2 > 1) // true
	t.Log(1 == 1 && 2 > 3) // false
	t.Log(1 == 1 || 2 > 3) // true
	t.Log(1 != 1 || 2 > 3) // false

	t.Log(!true)     // false
	t.Log(1 != 1)    // false
	t.Log(!(1 != 1)) // true

	// 如果 && 左边为假，则不判断右边
	// 如果 || 右边为真，则不判断右边
}

func TestRelationalOperator(t *testing.T) {
	// ----------------- 关系运算符 ---------------------
	t.Log(1 == 1)  // true
	t.Log(1 != -1) // true
	t.Log(1 < -1)  // false
	t.Log(1 > -1)  // true
	t.Log(1 >= -1) // true
	t.Log(1 <= -1) // false

	flag := 1 > 1
	t.Log(flag) // false

	flag2 := 1 == 1
	t.Log(flag2) // true
}

func TestBitOperator(t *testing.T) {
	// ---------------- 位运算符 -------------------
	var a int8 = 3
	var b int8 = -2

	// 左移
	t.Log(a << 1) // 6
	t.Log(b << 1) // -4

	// 右移
	t.Log(a >> 1) // 1
	// 负数右移最小值为 -1，操作原码
	t.Log(b >> 1) // -1

	// 负数按照补码计算: | & ^
	t.Log(a | b) // -1
	t.Log(a & b) // 2
	t.Log(a ^ b) // -3
	t.Log(a &^ b) // 1

	// 按位清零 ，右边为 1 左边无论是 0 还是 1 结果都是 0，右边为 0 左边是什么结果就是什么
	t.Log(1 &^ 0)
	t.Log(1 &^ 1)
	t.Log(0 &^ 0)
	t.Log(0 &^ 1)
}
