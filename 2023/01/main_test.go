package main

import (
	"testing"
)

func TestParseString(t *testing.T) {
	got := ParseString("jhsd3jhsd44hj8d")
	want := 38

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = ParseString("5")
	want = 55

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
} 

func TestIsByteNumber(t *testing.T) {

	got := IsByteNumber(byte('7'))
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = IsByteNumber(byte('a'))
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
