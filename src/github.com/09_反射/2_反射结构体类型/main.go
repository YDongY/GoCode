package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTestStruct(b interface{}) {
	// 1. 通过反射获取参数的　type kind 值
	rType := reflect.TypeOf(b)
	fmt.Println(rType) // main.Student

	// 2. 获取到 reflect.Value

	rValue := reflect.ValueOf(b)
	fmt.Printf("%v,%T\n", rValue, rValue.Type) // {tom 22},func() reflect.Type

	// 3. reflect.Value ==> interface{}
	iv := rValue.Interface()
	fmt.Printf("%v,%T\n", iv, iv) // {tom 22},main.Student ,但是 iv 无法获取字段

	// 4. 将 interface{} 通过断言转成结构体类型，获取结构体字段
	// 这里可以使用 switch 对所有类型进行断言判断
	fmt.Println(iv.(Student).Name) // tom


	// 获取变量对应的 kind
	fmt.Printf("kind = %v \n", rType.Kind())  // kind = struct
	fmt.Printf("kind = %v \n", rValue.Kind()) // kind = struct
}

func main() {
	reflectTestStruct(Student{Name: "tom", Age: 20})
}
