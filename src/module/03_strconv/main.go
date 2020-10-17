package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "false"
	b, _ := strconv.ParseBool(str)
	fmt.Printf("type %T b=%v \n", b, b) //type bool b=false

	n := "12345"
	var i int64
	i, _ = strconv.ParseInt(n, 10, 0)
	fmt.Printf("type %T ,i=%v \n", i, i)

}
