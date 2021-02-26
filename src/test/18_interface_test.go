package test

import (
	"fmt"
	"testing"
)

// 定义DataWriter接口
type DataWriter interface {
	WriteData(data interface{}) error
	// 能否写入
	CanWrite() bool
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

func (f *file) CanWrite() bool {
	panic("implement me")
}

// 实现DataWriter接口的WriteData方法
// 有些人可能会觉得，下面重复写接口的方法名、参数和返回值很麻烦
// 其实一般IDE都带有快速实现自动生成的功能，比如GoLand在定义结构体的类型名上按下Ctrl+I即可快速实现接口
//func (f *file) WriteDataX(data interface{}) error {
func (f *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

func TestInterface(t *testing.T) {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("data")
}

