package function

import (
	"fmt"
	"testing"
	"time"
)

// Cal1 有参数有返回值
func Cal1(a, b int) int {
	return a + b
}

// Cal2 多个返回值
func Cal2(a int, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

// Cal3 命名返回值
func Cal3(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func TestFuncDefine(t *testing.T) {
	t.Log(Cal1(1, 2))
	t.Log(Cal2(1, 2))
	t.Log(Cal3(1, 2))
}

func TestAnonymousFunc(t *testing.T) {
	// 匿名函数变量
	a := func(s string) {
		t.Logf(s)
	}
	a("匿名函数")
}

// timeSpent 接受函数，返回一个函数
func timeSpent(inner func(op int) int) func(op int) int {
	// inner 是函数类型
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestFuncType(t *testing.T) {
	slowFunc := func(op int) int {
		time.Sleep(time.Second * 1)
		return op
	}
	t.Log(timeSpent(slowFunc)(10))
}

// Sum 接受可变长参数
func Sum(ops ...int) int {
	sum := 0
	for _, i := range ops {
		sum += i
	}
	return sum
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5)) // 15
	t.Log(Sum(1, 2, 3, 4))    // 10
}
