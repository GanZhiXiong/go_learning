package test

import (
	"fmt"
	"testing"
)

type Animal interface {
	Say()
}

type Cat struct {
}

// Method redeclared '*Cat.Say'
//func (c Cat) Say() {
//	fmt.Println("meow")
//}

func (c *Cat) Say() {
	fmt.Println("meow")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("wang wang")
}

func TestInterfacePointer(t *testing.T) {
	/*
		cat
	*/
	c := Cat{}
	c.Say()

	c1 := &Cat{}
	c1.Say()

	// Cannot use 'Cat{}' (type Cat) as type Animal Type
	// does not implement 'Animal' as 'Say' method has a pointer receiver
	//var c2 Animal = Cat{}
	//c2.Say()

	var c3 Animal = &Cat{}
	c3.Say()

	/*
		dog
	*/
	d := Dog{}
	d.Say()

	d1 := &Dog{}
	d1.Say()

	var d2 Animal = Dog{}
	d2.Say()

	var d3 Animal = &Dog{}
	d3.Say()
}
