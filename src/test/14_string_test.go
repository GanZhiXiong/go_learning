package test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestString1(t *testing.T) {
	var s string
	t.Log("*"+s+"*", len(s))

	s = "hello"
	// Cannot assign to s[1]
	//s[1] = "a"
	t.Log(s, len(s))

	// 可以存储任何二进制数据
	s = "\xE4\xB8\xA5"
	// 输出：严 3
	t.Log(s, len(s))
}

func TestStringType(t *testing.T) {
	var s string
	s = "a"
	t.Log(s, &s)

	s1 := s
	t.Log(s1, &s1)
}

func TestStringLength(t *testing.T) {
	s := "Go中国"
	t.Log(s, len(s), len([]rune(s)))

	t.Log(utf8.RuneCountInString(s))
}

func TestGetChar(t *testing.T) {
	s := "Go中国"
	t.Log(s[0])
	t.Log(s[1])
	t.Log(s[2])
	t.Log(s[3])

	// 中文会输出乱码
	t.Log(s[0:1])
	t.Log(s[1:2])
	t.Log(s[2:3])
	t.Log(s[3:4])

	t.Log([]rune(s))
	t.Log(string([]rune(s)[0]))
	t.Log(string([]rune(s)[1]))
	t.Log(string([]rune(s)[2]))
	t.Log(string([]rune(s)[3]))
}

func TestUnicodeAndUTF8(t *testing.T) {
	var s string
	s = "中"
	t.Log(s, len(s))

	// 通过rune取出字符串里的Unicode，得到rune的切片
	c := []rune(s)
	t.Log(len(c))
	// %x表示用十六进制输出
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "Go中国"
	for _, c := range s {
		// [1]表示都是对应第一个参数
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStrings(t *testing.T) {
	s := "!a, b#, @c"
	parts := strings.Split(s, ",")
	for index, part := range parts {
		part = strings.Trim(part, "!@# ")
		parts[index] = part
		t.Log(index, part)
	}
	t.Log(strings.Join(parts, "-"))

	slice := []string{"a", "b", "c"}
	for index, part := range slice {
		slice[index] = part + fmt.Sprintf("%d", index)
		t.Log(index, part)
	}
	t.Log(slice)
}

func TestStrConv(t *testing.T) {
	i := 12
	t.Log(reflect.TypeOf(i))

	// 整形转字符串
	s := strconv.Itoa(i)
	t.Log(reflect.TypeOf(s))
	t.Log("0" + s)

	// 字符串转整形
	// 下面这么写会报错：Invalid operation: 10 + strconv.Atoi(s)
	// (mismatched types untyped int and (int, error))
	//t.Log(10 + strconv.Atoi(s))
	if ret, err := strconv.Atoi(s); err == nil {
		t.Log(10 + ret)
	}
}
