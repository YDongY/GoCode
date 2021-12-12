package _6_type_assert

import (
	"fmt"
	"testing"
)

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		index = index + 1
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数，类型是 bool 值是%v \n", index, x)
		case float32:
			fmt.Printf("第%v个参数，类型是 float32 值是%v \n", index, x)
		case float64:
			fmt.Printf("第%v个参数，类型是 float64 值是%v \n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数，类型是 int,int32,int64 值是%v \n", index, x)
		default:
			fmt.Printf("第%v个参数，类型不确定 值是%v \n", index, x)
		}
	}
}

func TestEmptyInterface(t *testing.T) {
	var n1 float32 = 1.1
	var n2 float64 = 12.3
	var n3 int32 = 30
	var name string = "tom"

	TypeJudge(n1, n2, n3, name)
}

type Point struct {
	x int
	y int
}

func TestAssertStruct(t *testing.T) {
	var a interface{} = Point{1, 2}
	b, ok := a.(Point)
	fmt.Printf("%T,%v,%v\n", b, b, ok) // _6_type_assert.Point,{1 2},true
}
