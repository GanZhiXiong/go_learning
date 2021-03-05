package test

import "testing"

type Binary uint64

func TestEface(t *testing.T) {
	b := Binary(200)
	t.Log(b)

	any := (interface{})(b)
	t.Log(any)
}
