package main

import "fmt"

type Data struct {
}

func dummy2() *Data {
	var c Data
	return &c
}

func dummy(b int) int {
	var c int
	b = c
	return c
}

func main() {
	var a int
	fmt.Println(a, dummy(0))
	fmt.Println(dummy2())
}

/**
$ go run -gcflags "-m -l" variable_life.go
# command-line-arguments
./variable_life.go:9:6: moved to heap: c
./variable_life.go:21:13: ... argument does not escape
./variable_life.go:21:13: a escapes to heap
./variable_life.go:21:22: dummy(0) escapes to heap
./variable_life.go:22:13: ... argument does not escape
0 0
&{}
*/
