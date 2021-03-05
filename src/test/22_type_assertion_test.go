package test

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

// 会报错
/*func funcName(a interface{}) string {
	// Cannot convert expression of type 'interface{}' to type 'string'
	return string(a)
}*/

func TestTypeAssertion(t *testing.T) {

	// x为非接口类型，所以编译不通过
	//var a = 1
	// Invalid type assertion: a.(int) (non-interface type int on left)
	//v, o := a.(int)
	//t.Log(v, o)

	var x interface{}
	x = 10
	t.Log(x, reflect.TypeOf(x))
	value, ok := x.(int)
	t.Log(value, ok)

	t.Log()

	x = 1.2
	t.Log(x, reflect.TypeOf(x))
	value, ok = x.(int)
	// 断言失败后，value为T类型，value为T类型的默认值
	t.Log(value, reflect.TypeOf(value), ok) // 0 false
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

	// 如果断言失败，省略了ok则会panic。
	t.Log(x.(int))
	value3 := x.(int)
	t.Log(value3)
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

func TestTypeAssertionSwitch(t *testing.T) {
	var a int
	a = 1
	getType(a)
}

type Html []interface{}

func TestTypeAssertionSwitch1(t *testing.T) {
	html := make(Html, 5)
	html[0] = "div"
	html[1] = "span"
	html[2] = []byte("script")
	html[3] = 1
	html[4] = 1.34
	t.Log(html, reflect.TypeOf(html))
	for index, element := range html {
		switch value := element.(type) {
		case string:
			t.Logf("html[%d] is a string, value is %s", index, value)
		case []byte:
			t.Logf("html[%d] is a []byte, value is %s", index, value)
		case int:
			t.Logf("html[%d] invalid type", index)
		default:
			t.Logf("html[%d] unknown type", index)
		}
	}
}

func TestTypeAssertionNil(t *testing.T) {
	var w io.Writer
	// 创建一个byte.Buffer指针类型的nil（即nil指针）
	w = (*bytes.Buffer)(nil)
	t.Log(w, reflect.TypeOf(w))
	t.Logf("%#v", w)
	x, ok := w.(io.Writer)
	t.Log(x, ok)

	w = nil
	t.Logf("%#v", w)
	x, ok = w.(io.Writer)
	t.Log(x, ok)
}

type A interface {
	Test()
}

type B struct{}

func (b B) Test() {
	panic("implement me")
}

type C struct{}

func (c *C) Test() {
	panic("implement me")
}

func TestTypeAssertion1(t *testing.T) {
	var a A = new(B)
	t.Logf("%#v", a)
	if false {
		v, ok := a.(*B)
		t.Log(v, ok) // ture
	} else {
		v, ok := a.(B)
		t.Log(v, ok) // false
	}

	var a1 A = new(C)
	t.Logf("%#v", a1)
	if true {
		v, ok := a1.(*C)
		t.Log(v, ok) // ture
	} else {
		// Impossible type assertion: 'C' does not implement 'A'
		// 不可能的类型断言:“C”没有实现“A”
		//v, ok := a1.(C)
		//t.Log(v, ok) // false
	}
}
