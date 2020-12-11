package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTestInt(b interface{}) {
	// 通过反射获取参数的　type kind 值
	rType := reflect.TypeOf(b)
	fmt.Println(rType) // int

	// 获取到 reflect.Value
	rValue := reflect.ValueOf(b)               // 100，实际不是数值100,无法进行计算
	fmt.Printf("%v,%T\n", rValue, rValue.Type) // 100,func() reflect.Type

	fmt.Printf("kind = %v \n", rType.Kind())  // kind = int
	fmt.Printf("kind = %v \n", rValue.Kind()) // kind = int

	// 获取真正的值
	fmt.Println(rValue.Int()) // 100，可用于计算

	// rValue 转 interface{}
	iv := rValue.Interface()

	// 将 interface{} 通过断言转成类型
	fmt.Println(iv.(int)) // 100

}

func reflectTestStruct(b interface{}) {
	// 通过反射获取参数的　type kind 值
	rType := reflect.TypeOf(b)
	fmt.Println(rType) // main.Student

	// 获取到 reflect.Value

	rValue := reflect.ValueOf(b)
	fmt.Printf("%v,%T\n", rValue, rValue.Type) // {tom 22},func() reflect.Type


	fmt.Printf("kind = %v \n", rType.Kind())  // kind = struct
	fmt.Printf("kind = %v \n", rValue.Kind()) // kind = struct


	iv := rValue.Interface()
	fmt.Printf("%v,%T\n", iv, iv) // {tom 22},main.Student ,但是iv 无法获取字段

	// 将 interface{} 通过断言转成类型

	fmt.Println(iv.(Student).Name) // tom

}

func main() {
	var num int = 100
	reflectTestInt(num)

	var stu = Student{
		Name: "tom",
		Age:  22,
	}

	reflectTestStruct(stu)

}
