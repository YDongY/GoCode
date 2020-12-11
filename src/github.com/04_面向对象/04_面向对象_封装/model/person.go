package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

// 工厂模式

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 120 {
		p.age = age
	} else {
		fmt.Println("年龄范围错误")
	}

}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary < 3500 {
		fmt.Println("薪水太少")
	} else {
		p.salary = salary
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
