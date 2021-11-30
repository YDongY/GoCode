package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) getName() string {
	return p.Name
}

func (p *Person) setName(name string) {

	// (*p).Name = name　等价于
	p.Name = name
}

func (p *Person) String() string { // 类似以Python 中的　__str__，方法名称必须是 String
	str := fmt.Sprintf("Person<Name:%s>", p.Name)

	return str
}

func main() {
	p := Person{"tom"}
	fmt.Println(p.getName())

	p.setName("jack")
	// (&p).setName("jack")

	fmt.Println(p.getName())

	fmt.Println(&p) // Person<Name:jack>
	// 默认调用 String　方法，类似于Python 中的__str__
}