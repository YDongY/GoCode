package main

import "fmt"

type Point struct {
	x int
	y int
}

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

func main() {
	var a interface{}
	var point = Point{1, 2}
	a = point
	b, ok := a.(Point) // 类型断言
	fmt.Println(b, ok)

	var n1 float32 = 1.1
	var n2 float64 = 12.3
	var n3 int32 = 30
	var name string = "tom"

	TypeJudge(n1, n2, n3, name)
}
