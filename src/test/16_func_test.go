package test

import "testing"

func sum(nums ...int) int {
	ret := 0
	for _, num := range nums {
		ret += num
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))

	nums := []int{1, 2, 3}
	t.Log(sum(nums...))

	//nums1 := [...]int{1, 2, 3}
	// Cannot use 'nums1' (type [3]int) as type []int
	//t.Log(sum(nums1...))

	nums1 := [...]int{1, 2, 3}
	t.Log(sum(nums1[:]...))

	//nums2 := [...]int{1, 2, 3}
	// Invalid use of ..., corresponding parameter is non-variadic
	// 无效使用…，对应的参数为非可变参数
	//t.Log(sum(4, nums2...))

	// 当然也可以不输入参数
	t.Log(sum())
	t.Log()
}

func TestVarParam1(t *testing.T) {
	result := []int{1, 2}
	data := []int{3, 4}
	// func append(slice []Type, elems ...Type) []Type
	result = append(result, data...)
	t.Log(result)

	result = append(result, 5, 6)
	t.Log(result)
}