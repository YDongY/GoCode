package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name  string // 名称首字母大写可以再其他包使用
	Age   int
	Color string
}

type Stu struct {
	name   string
	age    int
	scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // map
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var cat Cat

	fmt.Println(cat) // { 0 }

	cat.Name = "小白"
	cat.Age = 3
	cat.Color = "白"
	fmt.Println(cat) // {小白 3 白}
	fmt.Println(cat.Name)
	fmt.Println(cat.Age)
	fmt.Println(cat.Color)

	// struct 内存布局

	/*
		　　　　　　默认值

		　　　　　　　 0x000191
		            ------------                                ------------
			cat---> |    ""    | Name           cat.Name="小白"  |   "小白" |
					------------                                ------------
					|    0     | Age
					------------
					|    ""    | Color
					------------
	*/

	// 结构体声明

	var s1 Stu
	fmt.Println(s1) // { 0 [0 0 0 0 0] <nil> [] map[]}

	// 注意　：slice ，map ，指针，默认值都是 nil
	// 在使用之前需要 make

	s1.slice = make([]int, 10)
	s1.slice[0] = 100

	s1.map1 = make(map[string]string)

	s1.map1["key"] = "value"

	fmt.Println(s1) // { 0 [0 0 0 0 0] <nil> [100 0 0 0 0 0 0 0 0 0] map[key:value]}

	// 结构体值类型

	var cat1 Cat

	cat1.Name = "小白"

	fmt.Println(cat1) // {小白 0 }
	cat2 := cat1      // 这里相当于值拷贝，cat2 的修改不会影响cat1

	cat2.Age = 3
	cat2.Name = "小黑"

	fmt.Println(cat2) // {小黑 3 }
	fmt.Println(cat1) // {小白 0 }

	// 可以通过指针进行修改
	cat3 := &cat1

	fmt.Println(*cat3) // {小白 0 }

	(*cat3).Name = "小花" //等价 cat3.Name = "小花"

	fmt.Println("修改后的:", cat1) // {小花 0 }

	// 创建结构体方式
	// var cat Cat
	// var cat Cat = Cat{}
	// var cat *Cat = new (Cat) 或者　 cat := new (Cat)
	// var cat *Cat = &Cat{}

	//
	person := Person{"ydongy", 23}

	jsonPerson, _ := json.Marshal(person)

	fmt.Println(string(jsonPerson)) // {"Name":"ydongy","Age":23}

	// {"name":"ydongy","age":23} // 利用反射

	/*   加上tag


	     type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}


	*/

	// 创建结构体变量，指定字段值
	person2 := Person{
		Name: "jack",
		Age:  23,
	}
	fmt.Println(person2) // {jack 23}

	// 返回结构体指针

	person3 := &Person{
		Name: "tom",
		Age:  22,
	}
	// 等价于　person3 *Person = &Person{}
	fmt.Println(*person3)       // {tom 22}
	fmt.Printf("%p\n", person3) // 0xc00009a1a0
}
