package test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestStructInstantiation(t *testing.T) {
	/*
		方式一、基本的实例化形式
	*/
	// e0为结构体的实例
	// 结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化。
	var e0 Employee
	t.Logf("%p", &e0)
	e0.Id = "e0"
	e0.Name = "name"
	t.Log(e0)

	/*
		方式二、创建指针类型的结构体
	*/
	// 注意这里返回的是引用/指针，相当于 e := &Employee{}
	e1 := new(Employee)
	// 通过实例的指针即可访问成员，不需要使用->(c/c++中需要使用)
	e1.Id = "3"
	e1.Name = "e1"
	e1.Age = 23
	// 其实Go语言为了方便开发者访问结构体指针的成员变量，
	// 使用了语法糖（Syntactic sugar）技术，将 e1.Name 形式转换为 (*e1).Name。
	t.Log(e1, e1.Name, (*e1).Name)

	/*
		方式三、取结构体的地址实例化
	*/
	e2 := &Employee{}
	e2.Name = "e2"
	t.Log(e2)
}

func TestStructInit(t *testing.T) {
	/*
		使用“键值对”初始化结构体
	*/
	e := Employee{
		Id:   "0",
		Name: "Jason",
		Age:  18,
	}
	t.Log(e)

	e1 := Employee{Id: "2", Name: "Jacky"}
	t.Log(e1)

	/*
		使用多个值的列表初始化结构体
	*/
	e2 := Employee{"id", "name", 28}
	// Too few values 值太少
	//e2 := Employee{"id", "name"}
	t.Log(e2)
}

type People struct {
	name  string
	// 成员为结构体的非指针类型，将报错：Invalid recursive type 'People'（无效递归类型“People”）
	//child1 People
	child *People
}

func TestStructStruct(t *testing.T) {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	t.Log(relation, *relation)
}

func(e People) string(t *testing.T) string {
	t.Logf("People地址：%p，name地址：%p，child地址：%p", &e, &e.name, &e.child)
	return fmt.Sprintf("%s的%s是%s", e.child.name, e.child.name, e.name)
}

func(e *People) string1(t *testing.T) string {
	t.Logf("People地址：%p，name地址：%p，child地址：%p", e, &e.name, &e.child)
	return fmt.Sprintf("%s的%s是%s", e.child.name, e.child.name, e.name)
}

func TestFuncStructParam(t *testing.T) {
	relation := People{
		name:  "爷爷",
		child: &People{name: "爸爸"},
	}
	t.Log(relation)
	t.Logf("People地址：%p，name地址：%p，child地址：%p", &relation, &relation.name, &relation.child)
	//t.Logf("%p", unsafe.Pointer(&relation.child)) // 带0x开头
	//t.Logf("%x", unsafe.Pointer(&relation.child)) // 不带0x开头

	t.Log()
	t.Log(relation.string(t))
	t.Log(relation.string1(t))
}