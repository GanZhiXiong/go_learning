package test

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

// 使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。
type MyInt int64

func TestImplicit(t *testing.T) {
	var a = 1
	var b = 2
	a = b
	t.Log(a, b)
	t.Log("a type:", reflect.TypeOf(a))
	t.Logf("b type: %T", b)

	// 即使int32在int64范围内，还是会报错，因为Go不支持隐式类型转换
	var c int32 = 3
	var d int64 = 4
	// 报错：Cannot use 'c' (type int32) as type int64
	//d = c
	// Invalid operation: c + d (mismatched types int32 and int64)
	//var and = c + d
	d = int64(c)
	t.Log(c, d)
	t.Log("c type:", reflect.TypeOf(c))
	t.Logf("d type: %T", d)

	// 别名和原有类型也不能进行隐式类型转换
	var e MyInt
	var f int64
	f = d
	// Cannot use 'd' (type int64) as type MyInt
	//e = d
	//e = f
	e = MyInt(d)
	t.Log(e, f)
	t.Log("c type:", reflect.TypeOf(e))
	t.Logf("d type: %T", e)

	//var g int = 1
	//var h float32 = 2
	//var and = g + h
}

func TestDataRange(t *testing.T) {
	t.Log("int8 范围：", math.MinInt8, math.MaxInt8)
	t.Log("int16 范围：", math.MinInt16, math.MaxInt16)
	t.Log("int32 范围：", math.MinInt32, math.MaxInt32)
	t.Log("int64 范围：", math.MinInt64, math.MaxInt64)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 进行指针运算会报错：
	// Invalid operation: aPtr + 1 (mismatched types *int and untyped int)
	// 无效操作:aPtr + 1(不匹配的类型*int和无类型int)
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*", len(s))
	t.Logf("%T", s)
}

func TestL(t *testing.T) {
	cpu := runtime.GOARCH
	t.Log(cpu)
	int_size := strconv.IntSize;
	t.Log(int_size)
}