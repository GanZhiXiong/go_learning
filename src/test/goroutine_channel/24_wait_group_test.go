package goroutine_channel

import (
	"net/http"
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

func TestWaitGroup1(t *testing.T) {
	t.Log("start")

	var wg sync.WaitGroup

	var urls = []string{
		"https://www.github.com/",
		"https://www.baidu.com/",
		"https://www.google.com/",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			_, err := http.Get(url)
			t.Log(url, err)
		}(url)
	}

	wg.Wait()

	t.Log("over")
}
