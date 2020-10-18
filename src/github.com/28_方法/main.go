package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() { // 給 Person 类绑定一个方法，p　变量可以随意
	p.Name = "jack"     // 此处修改不会改变原 Person 的Name
	fmt.Println(p.Name) // jack
}

func (p Person) speak() {
	fmt.Printf("%s 是个好人 \n", p.Name)
}

func (p Person) cal() {
	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("计算结果:", sum)
}

func (p Person) cal2(n int) { // 接收参数
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	fmt.Println("计算结果:", sum)
}

func (p Person) cal3(n int, m int) int { // 接收参数,有返回值
	return n + m
}

func main() {
	p := Person{"tom"}
	p.test() //　调用方法，会把 p 传到 方法,属于值传递

	fmt.Println(p.Name) // tom

	p.speak()
	p.cal()
	p.cal2(10)
	res := p.cal3(10, 20)
	fmt.Println(res)
}
