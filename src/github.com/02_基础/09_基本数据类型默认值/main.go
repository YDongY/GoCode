package main

import (
	"fmt"
)

func defaultValue() {
	var str string               // ""
	var i int                    // 0
	var f float64                // 0.000000
	var b bool                   // false
	var c interface{}            // nil
	fmt.Println(str, i, f, b, c) // ""  0 0 false <nil>
}

func defaultType() {
	str := "hello"
	i := 10
	f := 3.14
	b := false

	fmt.Printf("%T %T %T %T \n", str, i, f, b) // string int float64 bool
}

func main() {
	defaultValue()
	defaultType()
}
