package goroutine_channel

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// wg用来等待程序结束
var wg sync.WaitGroup

func TestUnbufferedChan(t *testing.T) {
	// 创建一个无缓冲通道
	court := make(chan int)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("A", court)
	go player("B", court)

	// 发球
	court <- 1

	// 灯带游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	fmt.Println(name)

	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 这是由于rand是一种伪随机数，你可以通过设定不同seed（种子）的方法来解决这个问题
		rand.Seed(time.Now().UnixNano())
		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}

func TestUnbufferedChan1(t *testing.T) {
	c := make(chan string)

	// 等待一组线程的结束
	var wg sync.WaitGroup
	// 设定等待完成的协程数
	wg.Add(2)

	go func() {
		// 当前协程执行完成，就是把需要等待的协程数减1
		defer wg.Done()
		t.Log("下面的代码（发送者）会阻塞")
		c <- `foo`
		t.Log("发送者停止阻塞")
	}()

	go func() {
		// 当前协程执行完成，就是把需要等待的协程数减1
		defer wg.Done()

		time.Sleep(time.Second * 2)
		t.Log(`Message: ` + <-c)
		t.Log("接收者已经收到了")
	}()

	t.Log("等待协程执行完成")
	// 等待协程执行完成
	wg.Wait()

	t.Log("完成")
}

func TestUnbufferedChan2(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	// 阻塞主协程
	block := make(chan int)

	// counter
	go func() {
		for i := 0; ; i++ {
			// 限制下循环速度
			time.Sleep(2 * time.Second)

			t.Log("counter", "准备发送i")
			naturals <- i
			t.Log("counter", "已发送i")
		}
	}()

	// squarer
	go func() {
		for {
			t.Log("squarer", "准备接收naturals")
			i := <-naturals
			t.Log("squarer", "已接收naturals")

			t.Log("squarer", "准备发送squares")
			squares <- i * i
			t.Log("squarer", "已发送squares")
		}
	}()

	// printer
	go func() {
		for {
			t.Log("printer", "准备接收squares")
			t.Log(<-squares)
			t.Log("printer", "已接收squares")
		}
	}()

	block <- 0
}

func TestUnbufferedChan3(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	// 阻塞主协程
	block := make(chan int)

	// counter
	go func() {
		for i := 0; i < 3; i++ {
			// 限制下循环速度
			time.Sleep(2 * time.Second)

			t.Log("counter", "准备发送i")
			naturals <- i
			t.Log("counter", "已发送i")
		}
	}()

	// squarer
	go func() {
		for {
			t.Log("squarer", "准备接收naturals")
			i := <-naturals
			t.Log("squarer", "已接收naturals")

			t.Log("squarer", "准备发送squares")
			squares <- i * i
			t.Log("squarer", "已发送squares")
		}
	}()

	// printer
	go func() {
		for {
			t.Log("printer", "准备接收squares")
			t.Log(<-squares)
			t.Log("printer", "已接收squares")
		}
	}()

	block <- 0
}

func TestUnbufferedChan4(t *testing.T) {
	naturals := make(chan int)
	squares := make(chan int)
	// 阻塞主协程
	block := make(chan int)

	// counter
	go func() {
		for i := 0; i < 3; i++ {
			// 限制下循环速度
			time.Sleep(2 * time.Second)

			t.Log("counter", "准备发送i")
			naturals <- i
			t.Log("counter", "已发送i")
		}
		close(naturals)
	}()

	// squarer
	go func() {
		// 只有naturals关闭后，for range才会结束
		for i := range naturals {
			t.Log("squarer", "准备发送squares")
			squares <- i * i
			t.Log("squarer", "已发送squares")
		}
		close(squares)
	}()

	// printer
	go func() {
		for res := range squares {
			t.Log("printer", "准备接收squares")
			t.Log(res)
			t.Log("printer", "已接收squares")
		}
		close(block)
	}()

	t.Log("结束", <-block)
}
