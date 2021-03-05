package test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestTA(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	t.Logf("%#v", w)
	value, ok := w.(*os.File)
	t.Logf("%#v %v", value, ok)

	//value, ok = w.(*bytes.Buffer)
}

// 断言只能用于接口值变量，不能用于普通变量
/*func TestTA1(t *testing.T) {
	s := os.Stdout
	// Invalid type assertion: s.(*os.File) (non-interface type *File on left)
	x, ok := s.(*os.File)
}*/

// 如何证明断言后返回的第一个接收值是*os.File变量
func TestTA2(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	// 断言会成功，返回的value是*os.File类型，而不是io.Writer类型
	value, ok := w.(*os.File)
	t.Logf("%#v %v", value, ok)

	value.Write([]byte("hello"))
	value.Close()
}

// 断言的类型T除了是普通类型之外，还可以是接口类型
func TestTA3(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	// ReadWriteCloser是接口类型
	value, ok := w.(io.ReadWriteCloser)
	// 断言成功，说明File类型也实现了ReadWriteCloser接口
	t.Logf("%#v %T %v", value, value, ok)

	value1, ok := w.(io.Writer)
	t.Logf("%#v %T %v", value1, value1, ok)

	value.Write([]byte("hi"))
	value.Close()

	value1.Write([]byte("hi"))
	// 报错Unresolved reference 'Close'，因为io.Writer只有Write方法
	//value1.Close()
}

// 如果断言的是nil接口变量，那么无论被断言的类型是什么这个类型断言都会失败
func TestTANi(t *testing.T) {
	var w io.Writer
	w = nil
	value, ok := w.(io.Writer)
	t.Log(value, ok)

	var i interface{}
	value1, ok := i.(interface{})
	t.Log(value1, ok)
}

// 如果断言的是具体值为nil的接口变量，则不会失败
func TestTANil1(t *testing.T) {
	var w io.Writer
	// 创建一个byte.Buffer指针类型的nil（即nil指针）
	w = (*bytes.Buffer)(nil)
	t.Log(w)
	t.Log(reflect.TypeOf(w))
	t.Logf("%T", w)
	t.Logf("%#v", w)

	t.Log()

	value, ok := w.(io.Writer)
	t.Logf("%#v", value)
	t.Log(value, ok)
}

/*
	为了解决WriteHeader函数中Write函数的参数需要将string转[]byte带来的内存拷贝，
	在改进版本WriteHeader1函数中通过断言访问行为（通过断言判断这个变量有没有某个方法），来尽量避免调用Write函数
*/
// 假设这是一个往http响应添加响应头的函数
// w是一个响应流，向这个响应流写入的数据会发送到客户端
func WriteHeader(w io.Writer, contentType string) error {
	// [[byte()对字符串要进行类型转换，类型转换会直接开辟一个contentType相同大小的内存空间，并进行数据拷贝。
	// 并且拷贝出来的[]byte的空间会很快释放
	// 这样的内存分配方式会降低Web服务器的效率
	// 那有没有办法避免这种拷贝呢？WriteHeader1函数就可以解决这个问题
	if _, err := w.Write([]byte("Content-type: " + contentType)); err != nil {
		return err
	}
	return nil
}

func WriteHeader1(w io.Writer, contentType string) error {
	if _, err := writeString(w, "Content-type: "+contentType); err != nil {
		return err
	}
	return nil
}

/*
其实 *os.File 有一个 WriteString 方法，这个方法可以直接往一个文件句柄中写入字符串，
这个方法会避免去分配一个临时的拷贝。
不仅是 *os.File，像 *bytes.Buffer ，*bufio.Writer 这些类型都有这个方法。
可是，我们不能保证所有的io.Writer接口类型都有这个WriteString方法。
*/
func writeString(w io.Writer, s string) (n int, err error) {
	// 定义一个接口，这个接口用于判断w是否有WriteString方法
	type ws interface {
		WriteString(s string) (n int, err error)
	}

	// 通过断言判断w是否属于ws这个接口，从而判断w是否有WriteString方法
	if value, ok := w.(ws); ok {
		fmt.Println("Use function WriteString")
		return value.WriteString(s)
	}

	// 如果w没有WriteString方法，那就只能调用Write方法，就只能用[]byte转类型了
	return w.Write([]byte("Content-type: " + s))
}

func TestWriteHeader(t *testing.T) {
	w := os.Stdout
	WriteHeader(w, "text/html")
	t.Log()
	WriteHeader1(w, "text/html")
}

func TestTANoInterface(t *testing.T) {
	w := os.Stdout
	value, ok := (interface{})(w).(io.Writer)
	// Invalid type assertion: w.(io.Writer) (non-interface type *File on left)
	//value, ok := w.(io.Writer)
	t.Log(value, reflect.TypeOf(value), ok)
}

// 形参是x interface{}，当实参传入就会自动将实参从普通变量转换为接口值变量
func isInt(x interface{}) bool {
	_, ok := x.(int)
	return ok
}

func TestTANoInterface1(t *testing.T) {
	// 实参字面量（10）
	t.Log(isInt(10))

	num := 20
	// 实参是变量
	t.Log(isInt(num))
}
