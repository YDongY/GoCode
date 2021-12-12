package _struct

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   int // 名称首字母大写可以在其他包使用
	Age  int
	Name string
}

// 接受者为值类型，在实例方法被调用时，实例的成员会进行值赋值
func (e Employee) Print1() string {
	return fmt.Sprintf("%d-%s-%d", e.Id, e.Name, e.Age)
}

// 接受者为指针类型，可以避免在实例方法被调用时，内存拷贝
func (e *Employee) Print2() string {
	return fmt.Sprintf("%d-%s-%d", e.Id, e.Name, e.Age)
}

func TestStructInit(t *testing.T) {
	// 创建结构体返回值类型
	c1 := Employee{1, 24, "mark"}                // 1. 按顺序赋值
	c2 := Employee{Age: 24, Id: 1, Name: "mark"} // 2. 指定字段赋值
	var c3 Employee                              // 3. 先声明再赋值
	c3.Id = 1
	c3.Age = 24
	c3.Name = "mark"
	t.Log(c1, c2, c3)

	// 创建结构体返回指针类型
	pc1 := &Employee{1, 24, "mark"}
	pc2 := &Employee{Age: 24, Id: 1, Name: "mark"}
	pc3 := new(Employee)
	pc3.Id = 1
	pc3.Age = 24
	pc3.Name = "mark"
	t.Log(pc1, pc2, pc3)

	// 创建值类型结构体，调用接受者为值和指针类型
	e := Employee{1, 24, "mark"}
	t.Log(e.Print1())
	t.Log(e.Print2()) // e 是个值而非指针，带指针接收者的方法也能被直接调用 Go 会将语句 e.Print2() 解释为 (&e).Print2()

	// 创建指针类型结构体，调用接受者为值和指针类型
	pe := &Employee{1, 24, "mark"}
	t.Log(pe.Print1()) // 以值为接收者的方法被调用时，接收者既能为值又能为指针 Go 会将语句 e.Print1() 解释为 (*e).Print1()
	t.Log(pe.Print2())
}
