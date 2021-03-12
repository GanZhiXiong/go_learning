package goroutine_channel

import (
	"testing"
	"time"
)

func TestBufferedChan(t *testing.T) {
	ch := make(chan int, 3)

	t.Log(ch, len(ch), cap(ch))

	ch <- 1
	ch <- 2
	ch <- 3
	// 下面打开取消注释，则会报错：fatal error: all goroutines are asleep - deadlock!
	// 因为通道只能存储3个元素
	//ch <- 4

	t.Log(ch, len(ch), cap(ch))

	t.Log(<-ch)
	t.Log(<-ch)
	t.Log(<-ch)
	//  因为通道内的元素都已经取出，所以运行后报错：fatal error: all goroutines are asleep - deadlock!
	//t.Log(<-ch)

	t.Log(ch, len(ch), cap(ch))
}

func TestBufferedChan1(t *testing.T) {
	response := make(chan string, 3)

	// 同时向3个站点发起请求，request是一个阻塞的方法，需要等待服务的响应
	go func() { response <- request("mirror1") }()
	go func() { response <- request("mirror2") }()
	go func() { response <- request("mirror3") }()

	// main goroutine则接受最快的响应，其他响应不接收
	t.Log(<-response)
	t.Log(<-response)
	t.Log(<-response)
}

func request(mirror string) string {
	if mirror == "mirror1" {
		time.Sleep(5 * time.Second / 10) // 通过sleep模拟请求和响应在网络中传输的时间和服务器处理请求的时间
	} else if mirror == "mirror2" {
		time.Sleep(3 * time.Second / 10)
	} else if mirror == "mirror3" {
		time.Sleep(7 * time.Second / 10)
	}

	return "response from " + mirror
}
