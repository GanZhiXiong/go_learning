package test

import (
	"reflect"
	"testing"
)

func TestArrayInit(t *testing.T) {
	var a [3]int
	t.Log(a)
	a[0] = 1
	t.Log(a, reflect.TypeOf(a))

	// 一维数组
	b := [3]int{1, 2, 3}
	t.Log(b)

	// 二维数组
	c := [2][2]int{
		{1, 2},
		{3, 4},
	}
	t.Log(c)
}

func TestArrayEach(t *testing.T) {
	a := [3]int{1, 2, 3}
	for i := 0; i < len(a); i++ {
		t.Log(a[i])
	}

	// 那有没有类型Java或C#中的foreach: foreach(int i in a)，也是有的
	for index, e := range a {
		t.Log(index, e)
	}

	// 那我们可不可以省略index?
	for e := range a {
		// 通过输出结果可以看到e是索引
		t.Log(e)
	}

	// 好吧，那我把index写上，但是我不用它行不？这样肯定不行啊，会报错：Unused variable 'index'
	// Go中可以“_”表示占位，不使用它也不会报错
	for _, e := range a {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	t.Log(a, reflect.TypeOf(a))

	// 数组截取：a[开始索引(包含), 不包含(不包含)]
	b := a[1:2]
	t.Log(b, reflect.TypeOf(b)) // 2
	t.Log(b, len(b), cap(b))

	t.Log(a[1:3], reflect.TypeOf(a[1:3])) // 2 3
	//t.Log(a[1:]) // 2 3 4 5

	c := a[3:]
	t.Log(c, len(c), cap(c))

	t.Log(a[:3]) // 1 2 3
}
