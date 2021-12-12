package _8_reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeAddValue(t *testing.T) {
	var f int64 = 10

	// 获取类型，值
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f)) // int64 10

	// 根据类型获取值
	t.Log(reflect.ValueOf(f).Type()) // int64
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 3.14
	CheckType(f)
}

type Employee struct {
	ID   string
	Name string `json:"name"`
	Age  int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{
		ID:   "1",
		Name: "mark",
		Age:  24,
	}

	t.Log(reflect.TypeOf(e))                                 // *_8_reflect.Employee
	t.Log(reflect.TypeOf(e).Kind())                          // ptr
	t.Log(reflect.TypeOf(e).Elem().Name())                   // Employee
	t.Log(reflect.New(reflect.TypeOf(e).Elem()).Interface()) // &{  0}

	// 根据字段名获取值
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name")) //  Name: value(mark), Type(reflect.Value)
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed to get 'Name' field")
	} else {
		t.Log("Tag:json", nameField.Tag.Get("json")) // Tag:json name
	}

	// 根据方法名调用方法
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(25)})
	t.Log("Updated Age:", e) // Updated Age: &{1 mark 25}
}
