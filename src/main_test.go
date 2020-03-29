package main

import "testing"

func TestCalc(t *testing.T) {
	got := soma(3, 2)
	want := 5

	if got != want {
		t.Errorf("calc(3, 2) \n got: %v \n want: %v", got, want)
	}
}
