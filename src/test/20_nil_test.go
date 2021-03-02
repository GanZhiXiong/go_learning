package test

import (
	"testing"
	"unsafe"
)

func TestDefaultValue(t *testing.T) {
	var a bool
	t.Log(a, &a)

	var b int
	t.Log(b, &b)

	var c float64
	t.Log(c, &c)

	var d string
	t.Log(d, &d)

	var e [2]int
	t.Log(e, &e)

	var p *int
	t.Logf("%p", p)
	t.Log(p, &p)
	t.Log(p == nil)

	var f []int
	t.Log(f, &f)
	t.Log(f == nil)

	var g map[string]int
	t.Log(g, &g)
	t.Log(g == nil)

	var h interface{}
	t.Log(h, &h)
	t.Log(h == nil)
}

func TestNil(t *testing.T) {
	var nil string
	nil = "hi"
	t.Log(nil)
}

func TestNilCompare(t *testing.T) {
	var arr []int
	var arr1 []int
	var num *int
	var num1 *int
	t.Log(arr == nil)
	t.Log(arr1 == nil)
	t.Log(num == nil)
	t.Log(num1 == nil)

	// 相同类型的nil值可能无法比较
	t.Log(num == num1)
	// Invalid operation: arr == arr1 (operator == is not defined on []int)
	//t.Log(arr == arr1)

	// 不同类型的nil值不能比较
	// Invalid operation: arr == num (mismatched types []int and *int)
	//t.Log(arr == num)

	t.Log(unsafe.Pointer(&arr))
	t.Log(&num)
	t.Logf("%p\n", arr)
	t.Logf("%p", num)
}
