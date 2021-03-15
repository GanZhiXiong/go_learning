package goroutine_channel

import (
	"fmt"
	"sync"
	"test/goroutine_channel/bank"
	"testing"
	"time"
)

var deposit = 1

// 两个goroutine并发多次对存款变量进行修改
// 当循环次数为一亿的时候，就会出现每次打印最终余额不一致的问题
// 耗时大概0.23秒
func TestSharedVariablesConcurrentlyTest(t *testing.T) {
	waiting := make(chan struct{})

	go func() {
		// 循环一亿次
		for i := 0; i < 100000000; i++ {
			deposit++
		}
		waiting <- struct{}{}
	}()

	for i := 0; i < 100000000; i++ {
		deposit--
	}

	<-waiting
	t.Log(deposit)
}

// 多个并发体竞争同一个资源的时候，就会出现数据不一致的问题
// 所以采用一个goroutine来修改余额，使用channel进行协程间通信
// 耗时大概八九十秒
func TestSharedVariablesConcurrentlyTest1(t *testing.T) {
	waiting := make(chan struct{})

	go func() {
		for i := 0; i < 100000000; i++ {
			//t.Log("a")
			bank.Deposit <- 1 // 存一元
		}
		waiting <- struct{}{}
	}()

	for i := 0; i < 100000000; i++ {
		//t.Log("b")
		bank.Deposit <- -1 // 取一元
	}

	<-waiting
	t.Log(<-bank.Search)
}

// 如果相让资源在多个goroutine中修改呢？
// 那只需要做到同一时刻只允许一个goroutine修改该资源即可
// 那就得使用互斥锁
// 使用互斥锁
// 耗时大概3秒
func TestSharedVariablesConcurrentlyTest2(t *testing.T) {
	var amount int
	var mutex sync.Mutex

	defer func(t time.Time) {
		fmt.Println(float64(time.Now().UnixNano() - t.UnixNano()) / 1e9)
	}(time.Now())

	waiting := make(chan struct{})
	//mutex = &sync.Mutex{}

	go func(){
		for i:=0; i < 100000000; i++{
			//t.Log("a")
			mutex.Lock()
			amount++
			mutex.Unlock()
		}
		waiting <- struct{}{}
	}()

	for i:=0; i < 100000000; i++{
		//t.Log("b")
		mutex.Lock()
		amount--
		mutex.Unlock()
	}

	<-waiting
	fmt.Println(amount)		// 最后查看余额
}

// 使用读写锁
// 耗时大概6秒
func TestSharedVariablesConcurrentlyTest3(t *testing.T) {
	var amount int
	var mutex sync.RWMutex

	defer func(t time.Time) {
		fmt.Println(float64(time.Now().UnixNano() - t.UnixNano()) / 1e9)
	}(time.Now())

	waiting := make(chan struct{})
	//mutex = sync.RWMutex{}

	go func(){
		for i:=0; i < 100000000; i++{
			//t.Log("a")
			mutex.Lock()
			amount++
			mutex.Unlock()
		}
		waiting <- struct{}{}
	}()

	for i:=0; i < 100000000; i++{
		//t.Log("b")
		mutex.Lock()
		amount--
		mutex.Unlock()
	}

	<-waiting
	fmt.Println(amount)		// 最后查看余额
}

func TestMutex(t *testing.T) {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for i := 1000000; i > 0; i-- {
				count ++
			}
			//fmt.Println(count)
			t.Log(count)
		}()
	}

	wg.Wait()
	t.Log(count)
	//fmt.Scanf("\n")  //等待子线程全部结束
}

func TestMutex1(t *testing.T) {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)

	var countGuard sync.Mutex

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for i := 1000000; i > 0; i-- {
				countGuard.Lock()
				count ++
				countGuard.Unlock()
			}
			t.Log(count)
		}()
	}

	wg.Wait()
	t.Log(count)
	//fmt.Scanf("\n")  //等待子线程全部结束
}

func TestRWMutex(t *testing.T) {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)

	var countGuard sync.RWMutex

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for i := 1000000; i > 0; i-- {
				countGuard.Lock()
				count ++
				countGuard.Unlock()
			}
			t.Log(count)
		}()
	}

	wg.Wait()
	t.Log(count)
	//fmt.Scanf("\n")  //等待子线程全部结束
}

func TestRWMutex1(t *testing.T) {
	var count int
	var wg sync.WaitGroup
	wg.Add(2)

	var countGuard sync.RWMutex

	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for i := 1000000; i > 0; i-- {
				countGuard.RLock()
				count ++
				countGuard.RUnlock()
			}
			t.Log(count)
		}()
	}

	wg.Wait()
	t.Log(count)
	//fmt.Scanf("\n")  //等待子线程全部结束
}