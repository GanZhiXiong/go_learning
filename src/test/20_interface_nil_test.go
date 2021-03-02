package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterfaceNil(t *testing.T) {
	var a interface{}
	t.Log(a)
	t.Log(a == nil) // true

	a = nil
	t.Log(a)
	t.Log(a == nil) // true

	// cannot convert nil to type int
	//var b interface{} = (int)(nil)
	var b interface{} = (*int)(nil)
	t.Log(b)
	t.Log(b == nil) // false
	b = 1
	t.Log(b == nil) // false

	t.Log()

	// 接口指针类型变量，底层结构为{*interface{}, nil}，所以不为nil
	var c interface{} = (*interface{})(nil)
	c = 123
	t.Log(c)
	t.Log(c == nil) // false

	var d = (*interface{})(nil)
	// Cannot use '2' (type untyped int) as type *interface{}
	//d = 2
	// 能编译通过，但是会panic，这是因为d指向的是一个无效的内存地址
	//*d = 2
	t.Log(d)
	t.Log(d == nil) // true
}

func TestInterfaceNil1(t *testing.T) {
	// val的底层结构应该是(int64, 1)
	var val interface{} = int64(1)
	t.Log(reflect.TypeOf(val))
	// 因为字面量的整数在Go中默认类型为int，所以val的底层结构变成了(int, 2)
	val = 2
	t.Log(reflect.TypeOf(val))
}

// 定义一个结构体
type MyImplement struct{}

// 实现fmt包中的Stringer接口的String方法
func (m *MyImplement) String() string {
	return "hello"
}

// 虽然返回的Stringer接口的value为nil，但是type带有*MyImplement信息，所以和nil比较依然不相等
func GetStringer() fmt.Stringer {
	var s *MyImplement
	//println(s == nil)
	return s
}

func TestGetStringer(t *testing.T) {
	if GetStringer() == nil {
		t.Log("GetStringer() == nil")
	} else {
		t.Log("GetStringer() != nil")
	}
}

type data struct{}

func (d *data) Error() string { return "" }

// 这样做不严谨，因为返回的指针p被包装成error类型，所以返回的底层结构为(*data, nil)
// 返回值nil相比则不相等，这不是预期的结果。
func test() error {
	var p *data
	return p
}

// 因此得用下面这种方式，如果出错了，返回p；没出错，则返回nil。
func testRight() error {
	var p *data

	isTestError := false
	if isTestError {
		return p
	}
	return nil
}

func TestNilError(t *testing.T) {
	var err error = test()
	if err == nil {
		t.Log("err is nil, success")
	} else {
		t.Log("err is not nil, error")
	}

	err = testRight()
	if err == nil {
		t.Log("err is nil, success")
	} else {
		t.Log("err is not nil, error")
	}
}

func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	var rv reflect.Value
	if v, ok := value.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(value)
	}
	switch rv.Kind() {
	case reflect.Chan,
		reflect.Map,
		reflect.Slice,
		reflect.Func,
		reflect.Ptr,
		reflect.Interface,
		reflect.UnsafePointer:
		return rv.IsNil()
	}
	return false
}

func TestInterfaceNilCompare(t *testing.T) {
	var s fmt.Stringer
	t.Log(s, IsNil(s))

	var i interface{}
	t.Log(i, IsNil(i))

	g := GetStringer()
	t.Log(g, IsNil(g))

	test := test()
	t.Log(test, IsNil(test))
	test1 := testRight()
	t.Log(test1, IsNil(test1))
}
