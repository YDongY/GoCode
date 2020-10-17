package main

import "fmt"

func main() {
	// = += -= *= /=
	fmt.Println(".......")
	a := 10
	b := 20

	c := a
	a = b
	b = c
	fmt.Printf("%d,%d \n", a, b)

	n, m := 10, 20

	x := n + m
	n = x - n
	m = x - n

	fmt.Printf("%d,%d \n", n, m)
}
