package main

import (
	"testing"
)

func TestMathFunc(t *testing.T) {
	value, div := 10.0, 7.0
	got := modValue(value, div)
	res := 3.0
	if got != res {
		t.Errorf("modValue() = %f; want %f", got, res)
	}
}
