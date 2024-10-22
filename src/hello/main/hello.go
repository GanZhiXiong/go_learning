// 包，表明代码所在的模块（包）
package main

// 引入代码依赖
// IDE会自动导入包，所以可以直接使用fmt，而不需要写下面代码
import (
	"fmt"
	"time"
)

// 程序入口，功能实现
//func main() {
//	//x := 200
//	//var any interface{} = x
//	//fmt.Println(any)
//	//return
//
//	if len(os.Args) > 1 {
//		fmt.Println("Args[1]: ", os.Args[1])
//	}
//	fmt.Println("Start")
//	fmt.Println("Hello, World!")
//	fmt.Println("End")
//	os.Exit(1)
//}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		time.Sleep(time.Second)
	}
}

func main() {
	go running()

	// 接受命令行输入，不作任何事情
	var input string
	fmt.Scanln(&input)
}