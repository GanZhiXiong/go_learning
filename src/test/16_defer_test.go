package test

import (
	"os"
	"sync"
	"testing"
)

func TestDefer(t *testing.T) {
	t.Log("defer begin")
	// 将defer放入延迟调用栈
	defer t.Log(1)
	defer t.Log(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer t.Log(3)
	t.Log("defer end")
}

func TestPanicDefer(t *testing.T) {
	defer t.Log("宕机后要做的事情1")
	defer t.Log("宕机后要做的事情2")
	panic("宕机")
	t.Log("宕机后，我是不能被输出的")
}

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

// 根据键读取值
func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

// 根据键读取值，使用defer的写法更简洁
func readValueByDefer(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 对共享资源解锁
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

func TestFaxieLock(t *testing.T) {
	valueByKey["key"] = 1
	t.Log(readValue("key"))

	valueByKey["key"] = 2
	t.Log(readValueByDefer("key"))
}

// 根据文件名获取文件大小
func fileSize(fileName string) int64 {
	//	根据文件名打开文件，返回文件句柄和错误
	f, err := os.Open(fileName)
	//	如果打开时发生错误（如文件没找到、文件被占用等），返回文件大小为0
	if err != nil {
		return 0
	}

	//	获取文件状态信息
	info, err := f.Stat()
	//	如果获取信息时发生错误，关闭文件并返回文件大小0
	if err != nil {
		f.Close()
		return 0
	}

	//	取文件大小
	size := info.Size()
	//	关闭文件
	f.Close()

	return size
}

// 根据文件名获取文件大小，使用defer的写法更简洁，且不会忘记释放文件资源
func fileSizeByDefer(fileName string) int64 {
	//	根据文件名打开文件，返回文件句柄和错误
	f, err := os.Open(fileName)
	//	如果打开时发生错误（如文件没找到、文件被占用等），返回文件大小为0
	if err != nil {
		return 0
	}

	// 延迟调用Close
	defer f.Close()

	//	获取文件状态信息
	info, err := f.Stat()
	//	如果获取信息时发生错误，关闭文件并返回文件大小0
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}

	//	取文件大小
	size := info.Size()

	// defer机制触发, 调用Close关闭文件
	return size
}

func TestGetFileSize(t *testing.T) {
	t.Log(fileSize("16_defer_test.go"))
	t.Log(fileSizeByDefer("16_defer_test.go"))

}

//变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的
func testDefer1() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func TestDefer1(t *testing.T) {
	t.Log(testDefer1())
}

func testDefer2() int {
	var ret int
	defer func() {
		ret++
	}()
	return ret
}

func TestDefer2(t *testing.T) {
	t.Log(testDefer2())
}