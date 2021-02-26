package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func TestMax(t *testing.T) {
	var a = 100
	var b = 200
	var ret = max(a, b)

	// 注意：测试用例中使用t.Log和t.Logf代替fmt.Print和fmt.Printf，不然运行后会警告：No tests were run
	//fmt.Printf("最大值是：%d", ret)

	t.Logf("最大值是：%d", ret)
}

func returnMultiValues() (int, int) {
	// 这是由于rand是一种伪随机数，你可以通过设定不同seed（种子）的方法来解决这个问题
	rand.Seed(time.Now().UnixNano())
	// 随机数in [0,n)
	return rand.Intn(10), rand.Intn(20)
}

func TestReturnMultiValues(t *testing.T) {
	// 同时使用多个返回值
	a, b := returnMultiValues()
	t.Log(a, b)

	// 或者只是一个返回值
	c, _ := returnMultiValues()
	t.Log(c)
}

// 使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。
type MyFunc func(op int)int

// 参数和返回值可以使函数
func timeSpent(inner MyFunc) MyFunc {
// 和上面代码一样的
//func timeSpent(inner func(op int) int) func(op int) int {
	// 变量可以是函数
	// 下面代码会报错，因为函数定义是必须按照函数签名展开的
	//fun := MyFunc {
	fun := func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
	return fun
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestFunctionalProgramming(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
