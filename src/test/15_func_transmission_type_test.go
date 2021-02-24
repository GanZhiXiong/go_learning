package test

import (
	"fmt"
	"reflect"
	"testing"
)

func modifyDigital(originalValue int, newValue int, t *testing.T) {
	t.Logf("%p %p", &originalValue, &newValue)
	originalValue = newValue
	t.Logf("%p %p", &originalValue, &newValue)
}

func TestModifyDigital(t *testing.T) {
	originalValue, newValue := 1, 2
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)

	modifyDigital(originalValue, newValue, t)
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)
}

func modifyString(originalValue, newValue string, t *testing.T) {
	t.Logf("%p %p", &originalValue, &newValue)
	originalValue = newValue
	t.Logf("%p %p", &originalValue, &newValue)
}

func TestModifyString(t *testing.T) {
	originalValue, newValue := "a", "b"
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)

	modifyString(originalValue, newValue, t)
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)
}

func modifySlice(originalValue, newValue []string, t *testing.T) {
	t.Logf("%p %p", &originalValue, &newValue)
	originalValue = newValue
	t.Logf("%p %p", &originalValue, &newValue)
}

func modifySliceByPtr(originalValue, newValue *[]string, t *testing.T) {
	t.Logf("%p %p", &originalValue, &newValue)
	*originalValue = []string{"b"}
	t.Logf("%p %p", &originalValue, &newValue)
}

func TestModifySlice(t *testing.T) {
	originalValue, newValue := []string{"a"}, []string{"b"}
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)

	modifySlice(originalValue, newValue, t)
	t.Log(originalValue, newValue)
	t.Logf("%p %p", &originalValue, &newValue)

	modifySliceByPtr(&originalValue, &newValue, t)
	t.Log(originalValue)
	t.Logf("%p", &originalValue)
}

func modifyDigitalByPtr(originalValue *int, newValue *int, t *testing.T) {
	// 下面打印的是地址，是没有传入到该函数之前的地址
	t.Log(originalValue, newValue) // 0xc00011e238 0xc00011e240
	// 下面打印指针的地址，看样看到传入的指针的地址是新的地址，实际上就进行了指针的拷贝
	t.Logf("%p %p", &originalValue, &newValue) // 0xc000116028 0xc000116030
	// 拷贝的指针指向0xc00011e238 0xc00011e240，所以能改变变量的值
	*originalValue = *newValue
	t.Logf("%p %p", &originalValue, &newValue) // 0xc000116028 0xc000116030
}

func TestModifyDigitalByPtr(t *testing.T) {
	originalValue, newValue := 1, 2
	// 1 2
	t.Log(originalValue, newValue)
	// 0xc00011e238 0xc00011e240
	t.Logf("%p %p", &originalValue, &newValue)

	p := &originalValue
	p1 := &newValue
	t.Log(p, p1)
	t.Logf("%p %p", &p, &p1)

	modifyDigitalByPtr(&originalValue, &newValue, t)
	// 2 2
	t.Log(originalValue, newValue)
	// 0xc00011e238 0xc00011e240
	t.Logf("%p %p", &originalValue, &newValue)
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &p)
	p["张三"] = 20
}

func TestModifyMap(t *testing.T) {
	//persons := make(map[string]int)
	persons := map[string]int{}
	persons["张三"] = 19
	mp := &persons
	fmt.Printf("原始map的内存地址是：%p\n", mp)
	modifyMap(persons)
	fmt.Println("map值被修改了，新值为:", persons)
}

type Person struct {
	Name string
}

func modifyStruct(p Person) {
	fmt.Printf("函数里接收到Person的内存地址是：%p\n", &p)
	p.Name = "李四"
}

func modifyStructByPtr(p *Person) {
	fmt.Printf("%p\n", p)
	fmt.Printf("函数里接收到Person的内存地址是：%p\n", &p)
	p.Name = "李四"
}

func TestModifyStruct(t *testing.T) {
	p := Person{"张三"}
	fmt.Printf("原始Person的内存地址是：%p\n", &p)
	modifyStruct(p)
	fmt.Println(p)

	modifyStructByPtr(&p)
	fmt.Println(p)
}

func modifySlice1(ages []int, t *testing.T) {
	t.Logf("%p\n", ages)
	t.Logf("函数里接收到slice的内存地址是%p\n", &ages)
	ages[0] = 1
	ages = append(ages, 222)
	ages = []int{1, 2}
	ages[0] = 2
}

func TestModifySlice1(t *testing.T) {
	ages := []int{6, 6, 6}
	t.Log(ages)
	t.Logf("%p\n", ages)
	t.Logf("原始slice的内存地址是%p\n", &ages)
	modifySlice1(ages, t)
	t.Log(ages)
}

func TestSharedMemory(t *testing.T) {
	var a = 1
	t.Log(a)
	t.Logf("%p", &a)

	var b, c = &a, &a
	t.Log(reflect.TypeOf(a), reflect.TypeOf(b))
	t.Log(a, b, c)
	t.Log(a, *b, *c)
	t.Logf("%p, %p, %p", &a, &b, &c)

	*b = 2
	t.Log(a, b, c)
	t.Log(a, *b, *c)
	t.Logf("%p, %p, %p", &a, &b, &c)

	a = 3
	t.Log(a, b, c)
	t.Log(a, *b, *c)
	t.Logf("%p, %p, %p", &a, &b, &c)
	t.Log(&a, &b, &c)
}

// 通过源码可知make获取的其实是个指针，map只不过是它的别名。
// make是字面量的包装，屏蔽了指针的操作，让我们更方便的使用，从这个角度来讲，map是引用类型，但是引用类型和传引用是两码事。
// 所以即使函数参数没有传递map指针，也可以修改map的元素的值及长度
func Test1(t *testing.T) {
	//m := make(map[int]int)
	m := map[int]int{}
	m[0] = 1
	fmt.Println("修改前", len(m), m[0], m)
	test1(m)
	fmt.Println("修改后", len(m), m[0], m)
}
func test1(m map[int]int) {
	m[0] = 11
	m[1] = 2
	m = map[int]int{1: 2, 3: 4}
	fmt.Println("函数内", len(m), m[0], m)
}

// slice结构体里面有个指针(array unsafe.Pointer)指向数组，还有长度和cap，后面两个只是简单的int。
// 通过方法我们只能修改数组（因为是指针），但是不能修改slice的len和cap
func Test2(t *testing.T) {
	s := []int{1}
	fmt.Println("修改前", len(s), cap(s), s)
	test2(s)
	fmt.Println("修改后", len(s), cap(s), s)
}
func test2(s []int) {
	s[0] = 11
	s = append(s, 2)
	fmt.Println("函数内", len(s), cap(s), s)
}

// 这次因为扩容，连元素都不能修改了，由此可见此时仅仅是值传递。
func Test3(t *testing.T) {
	s := []int{1}
	fmt.Println("修改前", len(s), cap(s), s)
	test3(s)
	fmt.Println("修改后", len(s), cap(s), s)
}
func test3(s []int) {
	s = append(s, 2)
	s[0] = 11
	fmt.Println("函数内", len(s), cap(s), s)
}