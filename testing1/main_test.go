package main

import "testing"

func TestGetArr(t *testing.T) {
	a := [10]int{0, 2, 3, -1, 5, 6, 7, 8, 9, 10}
	if a != getArr() {
		t.Fail()
	}
}