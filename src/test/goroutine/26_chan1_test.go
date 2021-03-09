package goroutine

import (
	"fmt"
	"testing"
)

// 消费者
func printer(c chan int) {
	for true {
		data := <-c

		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	fmt.Println("接收方：收到，我这里也退出了循环！")
	c <- 0
	// 下面的代码可能会不执行，因为主协程结束，子协程就会终止运行
	//fmt.Println("收到，我这里也退出了循环！")
}

func TestChan11(t *testing.T) {
	c := make(chan int)
	go printer(c)

	// 生产者
	for i := 1; i <= 10; i++ {
		c <- i
	}

	t.Log("发送方：所有数据已发送完毕！")
	c <- 0

	<-c
	t.Log("发送方：好的，已收到你退出循环的消息！")
}
