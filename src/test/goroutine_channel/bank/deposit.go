package bank

var Deposit chan int // 存钱或取钱用到的channel
var Search chan int  // 查看余额用到的channel
var amount int       // 余额

func init() {
	Deposit = make(chan int)
	Search = make(chan int)

	// 必须开一个goroutine跑，否者会阻塞init方法，也会阻塞main可执行文件
	go func() {
		for {
			select {
			case money := <-Deposit: // 客户端存钱或取钱时走这个case
				amount += money
			case Search <- amount: // 客户端要查看余额时走这个case
			}
		}
	}()
}
