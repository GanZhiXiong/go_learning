package test

import "testing"

func TestLoop(t *testing.T) {
	t.Log("你可以这样写")
	i := 0
	for ; i < 5; i++ {
		t.Log(i)
	}

	t.Log("也可以这写")
	for i = 1; i < 5; i++ {
		t.Log(i)
	}

	t.Log("也可以这写")
	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	for {
		t.Log(n)
		n++
	}
}