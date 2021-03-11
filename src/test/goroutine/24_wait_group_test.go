package goroutine

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	const N = 5
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			t.Log(i)
		}(i)
	}

	t.Log("等待其他协程结束")
	wg.Wait()
	t.Log("其他协程已结束")
}
