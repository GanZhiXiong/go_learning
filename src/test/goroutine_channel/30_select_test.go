package goroutine_channel

import (
	"fmt"
	"testing"
	"time"
)

// 非阻塞的收发
func TestSelect(t *testing.T) {
	ch := make(chan int)
	select {
	case i := <-ch:
		t.Log(i)
	default:
		t.Log("default")
	}
}

// 随机执行
func TestSelect2(t *testing.T) {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		t.Log(time.Now())
		select {
		case <-ch:
			t.Log("case1")
		case <-ch:
			t.Log("case2")
		}
	}
}

func TestSelect3(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			t.Log(time.Now())
			select {
			case num := <-ch:
				t.Log("num = ", num)
			case <-time.After(2 * time.Second):
				t.Log("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	t.Log("程序结束")
}

func TestSelect4(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 100; i++ {
		select {
		case x := <-ch:
			t.Log(x)
		case ch <- i:
		}
	}
}

func TestSelect5(t *testing.T) {
	t.Log(time.Now())

	ch := make(chan int)
	go func(c chan int) {
		time.Sleep(time.Second * 3)
		ch <- 1
	}(ch)

	select {
	case v := <-ch:
		t.Log(time.Now())
		t.Log(v)
	case <-time.After(2 * time.Second): // 等待2秒
		t.Log(time.Now())
		t.Log("超时")
	}

	time.Sleep(time.Second * 5)
}

func TestSelect6(t *testing.T) {
	select {}
}

func TestSelect7(t *testing.T) {
	var ch chan int
	ch = make(chan int)

	block := make(chan bool)

	go func(c chan int) {
		c <- 100
	}(ch)

	for {
		t.Log(time.Now())
		select {
		// Receive may block because of 'nil' channel
		case <-ch:
			t.Log("ok")
			block <- true
		default:
			t.Log("ll")
		}
	}

	//time.Sleep(time.Second * 5)
	<-block
}

func TestSelect8(t *testing.T) {
	var ch chan int
	//ch = make(chan int)

	select {
	// Receive may block because of 'nil' channel
	case <-ch:
		t.Log("ok")
		//default:
		//	t.Log("ll")
	}

}

func TestSelect9(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)
}
func pump1(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println(i)
		ch <- i * 2
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		fmt.Println(i)
		ch <- i + 5
	}
}
func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

func TestSelect10(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
	default:
		t.Log("channel is full!")
	}
}