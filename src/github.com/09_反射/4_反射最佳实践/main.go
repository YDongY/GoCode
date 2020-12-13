package main

import (
	"fmt"
	"reflect"
)

/**

使用反射来遍历结构体字段、调用结构体方法、获取结构体标签的值

*/

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float64
	Sex   bool
}

func (p Person) Add(a int, b int) int {
	return a + b
}

func (p Person) Set(name string, age int, score float64, sex bool) {
	p.Name = name
	p.Age = age
	p.Score = score
	p.Sex = sex
}

func (p Person) Print() {
	fmt.Println("---start---")
	fmt.Println(p)
	fmt.Println("---end---")
}

func TestStruct(a interface{}) {
	// 获取到 reflect.Type 类型
	typ := reflect.TypeOf(a)
	// 获取到 reflect.Value 类型
	val := reflect.ValueOf(a)

	// 获取 a 对应的类别
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	// 获取该结构体一共多少字段
	num := val.NumField()
	fmt.Println("NumField:", num) // NumField: 4

	// 遍历结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d : 值＝ %v \n", i, val.Field(i))
		// 获取到 struct 标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d : Tag: %v\n", i, tagVal)
		}
	}

	// 获取结构体一共有多少方法
	numOfMethod := val.NumMethod()
	fmt.Println("NumMethod:", numOfMethod) // NumMethod: 2

	// var params []reflect.Value
	val.Method(1).Call(nil) // 调用第 2个方法，方法的排序是按照方法名 ASCII

	// 调用第一个方法，Add 带有参数
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10)) // 第一个参数
	params = append(params, reflect.ValueOf(20)) // 第二个参数
	res := val.Method(0).Call(params) // res 还是切片
	fmt.Println("res=", res[0].Int())

}

func main() {
	var p Person = Person{
		Name:  "jack",
		Age:   20,
		Sex:   true,
		Score: 99.9,
	}

	TestStruct(p)
}
