package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func test(i int) {
	fmt.Println(i)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 使用go关键字的方式，将函数放在一个新创建的goroutine协程上运行
		//go test(i)

		// 可能每次输出都是最后一个,
		// 因为启动goroutine的速度小于循环执行的速度，且闭包共享外部的变量i
		//go func() {
		//	fmt.Println(i)
		//}()

		// 函数参数都是值传递，可以将i传入匿名函数，解决上面这个问题
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	// TestGoroutine函数在一个单独的goroutine上运行
	// 但是当函数运行完成后，该函数中所有的子协程都会终止
	// 这里为了演示，所以加上个sleep
	// 注释该代码后，子协程不一定全部执行成功
	time.Sleep(time.Microsecond * 50)
}

func test1() {
	fmt.Println("子协程在运行")
	time.Sleep(time.Second)
	fmt.Println("子协程运行结束")
}

func TestGoroutine1(t *testing.T) {
	go test1()
	t.Log("主协程运行")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func TestSay(t *testing.T) {
	go say("world")
	say("hello")
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		time.Sleep(time.Second)
	}
}

func TestRunning(t *testing.T) {
	go running()

	// 接受命令行输入，不作任何事情
	var input string
	fmt.Scanln(&input)
}