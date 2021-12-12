package _3_interface

import (
	"fmt"
	"testing"
	"time"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() string {
	return `fmt.Println("Hello World")`
}

type PyProgrammer struct{}

func (p *PyProgrammer) WriteHelloWorld() string {
	return `print("Hello World")`
}

func TestProgrammer(t *testing.T) {
	g := new(GoProgrammer)
	t.Log(g.WriteHelloWorld())

	p := new(PyProgrammer)
	t.Log(p.WriteHelloWorld())
}

func TestInterfaceVar(t *testing.T) {
	// 接口变量
	var p Programmer = &GoProgrammer{}
	t.Logf("%T,%v", p, p) // *_3_interface.GoProgrammer,&{}
}

func TestInterfaceConversion(t *testing.T) {
	// 接口转换
	programmers := map[string]interface{}{
		"golang": new(GoProgrammer),
		"python": new(PyProgrammer),
	}

	for _, v := range programmers {
		programmer := v.(Programmer)
		t.Log(programmer.WriteHelloWorld())
	}
}

type FuncType func(op int) int // 自定义方法类型

func timeSpent(inner FuncType) FuncType { // 使用自定义方法类型
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func SlowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op * 2
}

func TestCustomType(t *testing.T) {
	t.Log(timeSpent(SlowFunc)(5))
}
