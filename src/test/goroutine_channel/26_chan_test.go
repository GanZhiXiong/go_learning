package goroutine_channel

import (
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch1 := make(chan int)
	t.Log(ch1)

	ch2 := make(chan interface{})
	t.Log(ch2)

	type a struct{}
	ch3 := make(chan *a)
	t.Log(ch3)
}

// fatal error: all goroutines are asleep - deadlock!
func TestChan1(t *testing.T) {
	ch := make(chan interface{})
	ch <- 0
	ch <- "hello"
}

// 通道接收数据方式3、阻塞接收任意数据，忽略接收的数据
// 使用通道做并发同步
func TestChan2(t *testing.T) {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		t.Log("start goroutine")

		// 通过通道通知TestChan2的goroutine
		ch <- 0

		t.Log("exit goroutine")
	}()

	t.Log("wait goroutine")

	// 等待匿名goroutine
	<-ch

	t.Log("all done")
}

// 通道接收数据方式4、循环接收
func TestChan3(t *testing.T) {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			t.Log("go func for", i)
			// 发送3到0之间的数值
			ch <- i

			// 每次发送完成时等待1s
			time.Sleep(time.Second)
		}
	}()

	// 遍历接收通道数据
	for data := range ch {
		t.Log(data)

		if data == 0 {
			break
		}
	}
}

// 测试题目1
func TestChanTest1(t *testing.T) {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			t.Log("go func for", i)
			// 发送3到0之间的数值
			ch <- i

			// 每次发送完成时等待1s
			time.Sleep(time.Second)
		}
	}()

	t.Log("wait")
	data := <-ch
	t.Log(data)

	t.Log("end")
}

// 测试题目2
func TestChanTest2(t *testing.T) {
	// 构建一个通道
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		for i := 3; i >= 0; i-- {
			t.Log("go func for", i)
			// 发送3到0之间的数值
			ch <- i

			// 每次发送完成时等待1s
			time.Sleep(time.Second)
		}
	}()

	t.Log("wait")
	data := <-ch
	t.Log(data)

	t.Log("time sleep start")
	time.Sleep(time.Second * 4)
	t.Log("time sleep end")
}
