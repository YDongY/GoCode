package main

import (
	"fmt"
	"reflect"
)

func reflectTestInt(b interface{}) {
	// 通过反射获取到传入变量的 type ,kind

	// 1. 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("%T %v \n", rType, rType) // *reflect.rtype int

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("%T %v\n", rVal, rVal) // reflect.Value ,100

	// 3. 反射获取变量值
	n := rVal.Int()
	fmt.Printf("%T %d \n", n, n) // int64 110

	// 4. rVal ==> interface
	iv := rVal.Interface()
	fmt.Printf("%T %d\n", iv, iv) // int 100

	// 5. 通过类型断言：interface ==> int
	n2 := iv.(int)

	fmt.Println(n2) // 100

}

func main() {
	var num int = 100
	reflectTestInt(num)
}
