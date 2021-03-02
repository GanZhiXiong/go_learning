package test

import (
	"fmt"
	"testing"
)

func TestTypeAssertion(t *testing.T) {
	var x interface{}
	x = 10
	t.Log(x)
	value, ok := x.(int)
	t.Log(value, ok)

	t.Log()

	x = 1.2
	t.Log(x)
	value, ok = x.(int)
	t.Log(value, ok)
	// Cannot assign float32 to value (type int) in multiple assignment
	//value, ok := x.(float32)
	value1, ok := x.(float32)
	t.Log(value1, ok)
	value2, ok := x.(float64)
	t.Log(value2, ok)
}

func TestTypeAssertionPanic(t *testing.T) {
	var x interface{}
	x = "Hello"
	value, ok := x.(string)
	t.Log(value, ok)

	value1 := x.(string)
	t.Log(value1)

	value2, ok := x.(int)
	t.Log(value2, ok)

	// 如果x不是int类型，省略了ok则会panic。
	value3 := x.(int)
	t.Log(value3)
}

func TestTypeAssertionSwitch(t *testing.T) {
	var a int
	a = 1
	getType(a)
}

func getType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	case float64:
		fmt.Println("is float")
	default:
		fmt.Println("unknown type")
		fmt.Println("test")

	}
}
