package main

import (
	"fmt"

	"github.com/30_面向对象_封装/model"
)

func main() {
	p := model.NewPerson("jack")
	fmt.Println(*p) // {jack 0 0}

	p.SetAge(18)
	p.SetSalary(5000)

	fmt.Println(*p) // {jack 18 5000}
}
