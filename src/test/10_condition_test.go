package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIf(t *testing.T) {
	n := -1
	//var n int
	if n < 0 {
		t.Log(n, "为负数")
	} else if n == 0 {
		t.Log(n, "为0")
	} else {
		t.Log(n, "为正数")
	}

	//// Non-bool '1' (type untyped int) used as condition
	//if 1 {
	//
	//}
}

func TestIf1(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}

	// 在后面我们回经常用到变量赋值的条件表达式，下例就是用到了Go的函数多返回值
	//if v, err := Fun(); err == nil {
	//
	//} else {
	//
	//}
}

func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("macOS")
	case "linux":
		t.Log("Linux")
	default:
		t.Logf("%s", os)
	}
}

func TestGetGrade(t *testing.T) {
	/* 定义局部变量 */
	var grade string
	var marks int = 50

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

func TestSwitchIfElse(t *testing.T) {
	n := 5
	switch {
	case n >= 0 && n <= 3:
		t.Log("0-3")
	case n >= 4 && n <=6:
		t.Log("4-6")
	case n >= 7 && n <= 9:
		t.Log("7-9")
	}
}