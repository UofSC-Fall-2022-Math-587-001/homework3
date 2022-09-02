package main

import (
	"testing"
)

func Test1(t *testing.T) {
	got := gcd(527,1258)[0]
	want := 17

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test2(t *testing.T) {
	got := gcd(228,1056)[0]
	want := 12

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test3(t *testing.T) {
	got := gcd(163961,167181)[0]
	want := 7 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test4(t *testing.T) {
	got := gcd(3892394,238847)[0]
	want := 1 

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

