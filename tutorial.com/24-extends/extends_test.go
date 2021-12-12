package _4_extends

import (
	"fmt"
	"testing"
)

// ------------------- 公共　-----------------------
// 公共
type Student struct {
	Name string
	Age  int
}

// 公共 方法
func (stu *Student) ShowInfo() {
	fmt.Println(stu.Name, stu.Age)
}

func (stu *Student) SayHello() {
	fmt.Println("Student say Hello")
}

// ------------------- 小学生 ----------------------

type Pupil struct {
	Student // 组合
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

// ------------------ 大学生 -----------------------

type Graduate struct {
	*Student // 组合
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}

// -------------------- 多个匿名结构体 --------------------------

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

// ------------------ 命名结构体 -------------------------------

type D struct {
	a A // 命名结构体
}

func TestExtends(t *testing.T) {
	// ------------------------ 小学生 ---------------------
	p := &Pupil{}
	p.Student.Name = "jack"
	p.Student.Age = 10
	p.Student.ShowInfo() // 也可以直接 p.ShowInfo()
	p.SayHello()
	p.testing()

	p = &Pupil{
		Student{
			Name: "jack",
			Age:  10,
		},
	}

	// ------------------ 大学生 -----------------------
	g := &Graduate{
		&Student{
			Name: "mark",
			Age:  24,
		},
	}

	g.ShowInfo()
	g.SayHello()
	g.testing()

	/**
	  注意：
	  	1. 当结构体和组合的匿名结构体拥有相同的字段和方法，优先调用自己的
	  	2. 当结构体和组合的匿名结构体拥有相同的字段和方法，如果想要调用组合结构体的字段和方法需要通过**显式调用匿名结构体**
	  	3. 当多个组合匿名结构体存在相同字段，而自身结构体没有，需要通过**显式调用某个匿名结构体**
	*/
}
