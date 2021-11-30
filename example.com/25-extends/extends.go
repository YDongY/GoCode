package main

import "fmt"

// ------------------- 公共　-----------------------
// 公共
type Student struct {
	Name  string
	Age   int
	Score int
}

// 公共 方法
func (stu *Student) ShowInfo() {
	fmt.Println(stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (stu *Student) SayHello() {
	fmt.Println("Student say Hello")
}

// ------------------- 小学生 ----------------------

type Pupil struct {
	Student // 继承
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试...")
}

// ------------------ 大学生 -----------------------

type Graduate struct {
	Student // 继承
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试...")
}

// ------------------- 其他 ----------------------
type SSS struct {
	Student // 继承
	Name    string
}

func (s *SSS) SayHello() { //
	fmt.Println("SSS say Hello")
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

// ------------------ 组合 -------------------------------

type D struct {
	a A // 有名结构体
}

// ------------------------------------------------------
type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

// ------------------------------------------------------

func main() {

	// 方式一
	p := &Pupil{}
	p.Student.Name = "jack"
	p.Student.Age = 10

	p.testing()
	p.Student.SetScore(100)
	p.Student.ShowInfo()

	g := &Graduate{}
	g.Student.Name = "tom"
	g.Student.Age = 22

	g.testing()
	g.Student.SetScore(100)
	g.Student.ShowInfo()

	// 方式二，简化
	p.Name = "jack"
	p.Age = 10

	p.testing()
	p.SetScore(100)
	p.ShowInfo()

	// 测试继承就近原则
	s := SSS{
		Name: "ydongy",
	}

	fmt.Println(s.Name)         // ydongy
	fmt.Println(s.Student.Name) // ""
	s.SayHello()                // SSS say Hello
	s.Student.SayHello()        // Student say Hello

	// 当本结构体和继承的匿名结构体拥有相同的字段和方法，优先查找自己的
	// 如果自身存在则不调用继承类的方法或字段
	// 如果想要调用继承类的字段和方法通过　显式调用匿名类

	// 当嵌入匿名结构体,其匿名结构存在相同字段，而自身结构体没有
	// 则无法直接调用，必须通过显式调用匿名结构名称

	// var c C
	// c.Name = "tom" // error
	// c.A.Name = "tom" // ok
	// c.B.Name = "tom" // ok

	// 有名结构体
	var d D
	d.a.Name = "tom"

	// 结构体赋值
	tv := TV{
		Goods{
			Name:  "电视机",
			Price: 5000,
		},
		Brand{
			Name:    "海尔",
			Address: "北京",
		},
	}

	fmt.Println(tv) // {{电视机 5000} {海尔 北京}}

	tv2 := TV2{
		&Goods{
			Name:  "电视机",
			Price: 5000,
		},
		&Brand{
			Name:    "海尔",
			Address: "北京",
		},
	}
	fmt.Println(tv2)                    // {0xc000114000 0xc000114020}
	fmt.Println(*tv2.Goods, *tv2.Brand) // {电视机 5000} {海尔 北京}
}