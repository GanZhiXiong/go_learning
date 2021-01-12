// 包，表明代码所在的模块（包）
package main

// 引入代码依赖
// IDE会自动导入包，所以可以直接使用fmt，而不需要写下面代码
import (
	"fmt"
	"os"
)

// 程序入口，功能实现
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Args[1]: ", os.Args[1])
	}
	fmt.Println("Start")
	fmt.Println("Hello, World!")
	fmt.Println("End")
	os.Exit(1)
}
