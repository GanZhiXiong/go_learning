package test

import (
	"testing"
)

func TestVar(t *testing.T) {
	// 和JavaScript一样使用var声明变量
	var v int
	t.Log(v)

	// 声明并初始胡一个变量
	var a int = 1
	t.Log(a)

	var a1 = 1
	t.Log(a1)

	b := 2
	t.Log(b)

	// 定义多个变量
	var (
		c int = 3
		d int = 4
	)
	t.Log(c, d)

	var (
		c1 int
		d1 string
	)
	t.Log(c1, d1)

	var e, f = 5, "f"
	t.Log(e, f)
}

func TestFibList(t *testing.T) {
	var a = 1
	var b = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	t.Log(a, b)

	tmp := a
	a = b
	b = tmp
	t.Log(a, b)

	a, b = b, a
	t.Log(a, b)
}