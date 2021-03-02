package test

import (
	"fmt"
	"reflect"
	"testing"
)

//func Print(v interface{}) {
//	fmt.Println(v)
//	println(v)
//	fmt.Printf("%T\n", v)
//	fmt.Print(reflect.TypeOf(v))
//}
//
//func TestEmptyInterface(t *testing.T) {
//	type Test struct {
//	}
//	v := Test{}
//	t.Log(reflect.TypeOf(v))
//	fmt.Printf("%T\n", v)
//
//	Print(v)
//}

func TestEmptyInterface(t *testing.T) {
	var a interface{} = 1
	t.Log(a, reflect.TypeOf(a))

	a = 1.34
	t.Log(a, reflect.TypeOf(a))

	a = "hello"
	t.Log(a, reflect.TypeOf(a))

	a = [2]int{1, 2}
	t.Log(a, reflect.TypeOf(a))
}

func TestEmptyInterfaceSlice(t *testing.T) {
	var dataSlice = []int{1, 2}
	t.Log(dataSlice)
	/*
		错误的写法
	*/
	// Cannot use 'dataSlice' (type []int) as type []interface{}
	//var interfaceSlice []interface{} = dataSlice
	//t.Log(interfaceSlice)

	/*
		正确的写法
	*/
	var interfaceSlice = make([]interface{}, len(dataSlice))
	for i, value := range dataSlice {
		interfaceSlice[i] = value
	}
	t.Log(interfaceSlice)
}

func echoArray(a interface{}) {
	fmt.Printf("%T\n", a)
	println(reflect.TypeOf(a))

	// 下面代码a会报错：Cannot range over 'a' (type interface{})
	// interface{}可以向函数传递任意类型的变量，但对于函数内部，该变量仍为interface{}类型，而不是[]int类型。
	//for _, v := range a{
	//	fmt.Print(v, " ")
	//}

	// 使用断言实现类型转换，将interface{}转换为slice类型。
	b, ok := a.([]int)
	if ok {
		for _, i := range b {
			fmt.Println(i)
		}
	}
}

func TestEmptyInterface1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	t.Log(reflect.TypeOf(a))
	echoArray(a)
}
