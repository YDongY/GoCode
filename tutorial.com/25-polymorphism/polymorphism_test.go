package _5_polymorphism

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct{}

func (p *GoProgrammer) WriteHelloWorld() string {
	return `fmt.Println("Hello World")`
}

type PyProgrammer struct{}

func (p *PyProgrammer) WriteHelloWorld() string {
	return `print("Hello World")`
}

func writeFirstProgrammer(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	g := new(GoProgrammer)
	p := new(PyProgrammer)
	writeFirstProgrammer(g)
	writeFirstProgrammer(p)
}
