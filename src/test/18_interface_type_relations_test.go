package test

import (
	"fmt"
	"io"
	"testing"
)

/*
一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。
*/
type Socket struct {
}

// Socket 结构的 Write() 方法实现了 io.Writer 接口：
func (s *Socket) Write(p []byte) (n int, err error) {
	return 0, nil
}

// Socket 结构也实现了 io.Closer 接口：
func (s *Socket) Close() error {
	return nil
}

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	writer.Write(nil)
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

func TestInterfaceTypeRelations(t *testing.T) {
	s := new(Socket)
	// usingWriter() 和 usingCloser() 完全独立，互相不知道对方的存在，
	// 也不知道自己使用的接口是 Socket 实现的。
	usingWriter(s)
	usingCloser(s)
	t.Log(s)
}

/*
多个类型可以实现相同的接口
*/
type Service interface {
	Start()  // 启动服务
	Log(string) // 日志输出
}

// 日志器
type Logger struct {
}

// 实现Service接口的Log方法
func (l *Logger) Log(s string) {
	fmt.Println(s)
}

// 游戏服务
type GameService struct {
	// 选择将 Logger 嵌入到 GameService 能最大程度地避免代码冗余，简化代码结构。
	Logger
}

func (g *GameService) Start() {
	fmt.Println("启动服务成功")
}

func TestGameService(t *testing.T) {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
}



