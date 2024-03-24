package main

import "testing"

func TestSum(t *testing.T) {
	if s := Sum(1, 2); s != 3 {
		t.Errorf("got %v, expect 3", s)
	}
}
