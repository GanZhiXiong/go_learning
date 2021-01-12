package test

import "testing"

func TestIncrementDecrement(t *testing.T) {
	a := 1
	a++
	// '++' unexpected
	//++a
	t.Log(a)

	b := 1
	b++
	// '--' unexpected
	//--b
	t.Log(b)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 3, 2, 4}
	d := [...]int{1, 2, 3, 4}

	// Invalid operation: a==b (mismatched types [4]int and [5]int)
	//t.Log(a==b)
	// false
	t.Log(a==c)
	// true
	t.Log(a==d)
}