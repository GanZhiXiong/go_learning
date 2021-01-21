package test

import (
	"reflect"
	"sort"
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1, len(m1))

	m2 := map[int]int{}
	t.Log(m2, len(m2))
	m2[0] = 0
	m2[5] = 5
	t.Log(m2, len(m2))

	m3 := make(map[string]string, 10)
	// 不能查看Map的Capacity，否者会报错：Invalid argument for cap
	//t.Log(m3, len(m3), cap(m3))
	t.Log(m3, len(m3))
	m3["a"] = "1"
	m3["b"] = "2"
	m3["c"] = "3"
	m3["d"] = "4"
	t.Log(m3, len(m3))
	t.Log(m3["d"])
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[0]) // 0这个key不存在，取int的默认值0为value
	m1[2] = 0
	t.Log(m1[2])

	m2 := map[string]string{}
	t.Log(m2["a"]) // a这个key不存在，取string的默认值空字符串为value
	m2["a"] = "1"
	t.Log(m2["a"])

	// 这样虽然可以避免在取一个不存在的key时，不会出现异常，但是不能知道这个key到底存不存在
	key := 3
	m1[key] = 31
	if v, exists := m1[3]; exists {
		t.Logf("Key %d's value is %d", key, v)
	} else {
		t.Logf("Key %d is not existing", key)
	}
}

func TestTravelMap(t *testing.T) {
	t.Log("Map遍历")
	m3 := make(map[string]string, 10)
	t.Log(reflect.TypeOf(m3))
	m3["a"] = "1"
	m3["b"] = "2"
	m3["c"] = "3"
	m3["d"] = "4"
	t.Log(m3)
	for key, value := range m3 {
		t.Log(key, value)
	}

	t.Log("Array遍历")
	array := [...]string{"a", "b", "c"}
	t.Log(reflect.TypeOf(array))
	for index, element := range array {
		t.Log(index, element)
	}

	t.Log("Slice遍历")
	slice := []string{"a", "b", "c"}
	t.Log(reflect.TypeOf(slice))
	for index, element := range slice {
		t.Log(index, element)
	}
}

func TestDeleteKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2])
	printKey3Exists(&m, t)
	m[3] = 0
	printKey3Exists(&m, t)
	delete(m, 3)
	printKey3Exists(&m, t)
}

func printKey3Exists(m *map[int]int, t *testing.T) {
	t.Log(m)

	if v, ok := (*m)[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestSortByKey(t *testing.T) {
	m1 := map[string]int{"aa": 1, "zz": 2, "bb": 3, "yy": 4, "cc": 5, "xx": 6}
	s1 := make([]string, 0, len(m1))
	for k, _ := range m1 {
		s1 = append(s1, k)
	}
	sort.Strings(s1)
	for _, k := range s1 {
		t.Log(k, m1[k])
	}
}

func TestSortByValue(t *testing.T) {
	m1 := map[string]int{"aa": 3, "zz": 1, "bb": 5, "yy": 2, "cc": 4, "xx": 6}
	valueSlice := make([]int, 0, len(m1))
	for _, value := range m1 {
		valueSlice = append(valueSlice, value)
	}
	sort.Ints(valueSlice)
	for _, element := range valueSlice {
		for key, value := range m1 {
			if element == value {
				t.Log(key, value)
			}
		}
	}
}
