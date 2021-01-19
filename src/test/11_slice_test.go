package test

import (
	"reflect"
	"testing"
)

func TestSliceInit(t *testing.T) {
	// 声明
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0), reflect.TypeOf(s0))
	t.Log("\r")

	// 声明并初始化
	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))
	s1 = append(s1, 5)
	t.Log(s1, len(s1), cap(s1))
	s1[0] = 0
	t.Log(s1, len(s1), cap(s1))
	t.Log("\r")

	// 使用make创建
	// 如果len大于cap则会报错：larger than cap in make([]int)
	//s2 := make([]int, 3, 0)
	s2 := make([]int, 3, 3)
	t.Log(s2, len(s2), cap(s2))
	// panic: runtime error: index out of range [3] with length 3 [recovered]
	//	panic: runtime error: index out of range [3] with length 3
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		t.Log(s2, len(s2), cap(s2))
	}
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}

func TestSliceCapGrowthRules(t *testing.T) {
	s := make([]int, 0)
	n := 0
	for n < 1500 {
		t.Log(len(s), cap(s))
		s = append(s, n)
		n++
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Ja", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(year, len(year), cap(year))

	summer := year[4:7]
	// 为什么summer的capacity是8呢？
	// 虽然截取到索引6为止，但是summer是指向连续的存储空间year，也就是从year索引4开始到year最后一个元素的这一段连续空间
	// 也就是从Jun到Dec，总共8个，capacity也就是8。
	// 如果还觉得很绕，可以先这样记，后面我也会讲到，到时候进一步消化。
	t.Log(summer, len(summer), cap(summer)) // [May Jun Jul] 3 8

	summer[0] = "Unknown"
	t.Log(summer, len(summer), cap(summer))
	t.Log(year, len(year), cap(year))
}

func TestSliceComparing(t *testing.T) {
	//a := []int{1, 2, 3}
	//b := []int{1, 2, 3}
	// 语法报错：Invalid operation: a == b (operator == is not defined on []int)
	// 编译时报错：invalid operation: a == b (slice can only be compared to nil)
	//if a == b {
	//	t.Log("equal")
	//}
}