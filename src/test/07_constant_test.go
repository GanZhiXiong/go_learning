package test

import (
	"fmt"
	"testing"
)

const (
	Monday = 2 + iota
	Tuesday
	Wednesday
)

const (
	Readable   = 1 << iota // 0001 可读
	Writable               // 0010 可写
	Executable             // 0100 可执行
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func TestConstant(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstant1(t *testing.T) {
	t.Log(Readable, Writable, Executable)
	//a := 7 // 7%2=1 3%2=1 1%2=1 0111
	a := 1 // 1%2=1 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 下面演示09讲中的按位置零运算符
	a = 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Readable
	a = a &^ Writable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 按位运算符其实就是运算符优先级省略掉括号的写法而已
	b := Readable | Executable
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
	t.Log(b&^Readable == Executable)
	t.Log(b&^Executable == Readable)
	t.Log(b&(^Readable) == Executable)
	t.Log(b&(^Executable) == Readable)
}

func TestConstant2(t *testing.T) {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}

func TestConstant3(t *testing.T) {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
