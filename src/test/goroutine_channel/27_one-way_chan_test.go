package goroutine_channel

import (
	"fmt"
	"testing"
	"time"
)

func TestOneWayChan(t *testing.T) {
	ch := make(chan int)
	// 声明一个只能写入数据的通道，并赋值为ch
	var chSendOnly chan<- int = ch
	t.Log(chSendOnly)
	chSendOnly <- 0
	// Invalid operation: <-chSendOnly (receive from send-only type chan<- int)
	//<-chSendOnly

	// 声明一个只能读取数据的通道，并赋值为ch
	var chReadOnly <-chan int = ch
	t.Log(chReadOnly)
	// Invalid operation: chReadOnly <- 0 (send to receive-only type <-chan int)
	//chReadOnly <- 0
	<-chReadOnly
}

func TestOneWayChanMake(t *testing.T) {
	// 一个不能写入数据只能读取的通道是毫无意义的，因为它是空的
	ch := make(<-chan int)
	var chReadOnly <-chan int = ch
	<-ch
	<-chReadOnly
}

func TestOneWayChanTime(t *testing.T) {
	timer := time.NewTimer(time.Second)
	t.Log(timer)
}

func TestOneWayChanClose(t *testing.T) {
	ch := make(<-chan int)
	t.Log(ch)
	// 只读通道不能关闭，否者会报错，Cannot use 'ch' (type <-chan int) as type chan<- Type
	//close(ch)

	ch1 := make(chan<- int)
	close(ch1)
}

func TestChanClose(t *testing.T) {
	ch := make(chan int)
	close(ch)

	t.Logf("ptr: %p cap: %d len: %d", ch, cap(ch), len(ch))

	ch <- 1
}

func TestChanClose1(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 0
	ch <- 1

	close(ch)

	t.Logf("ptr: %p cap: %d len: %d", ch, cap(ch), len(ch))

	for i := 0; i < cap(ch)+2; i++ {
		v, ok := <- ch
		t.Log(v, ok)
	}
}

func TestChanClose2(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func TestChanClose3(t *testing.T) {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func TestChanClose4(t *testing.T) {
	var ch chan int
	t.Log(ch)
	close(ch)
}