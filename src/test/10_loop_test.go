package test

import "testing"

func TestLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(i)
	}
}