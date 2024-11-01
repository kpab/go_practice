package main

import "testing"

var Debug bool = false

func TestIsObe(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("Skipさせる")
	}
	v := IsOne(i)

	if !v {
		t.Errorf("%v ! = %v", i, 1)
	}
}
