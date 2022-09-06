package main

import (
	"testing"
	"sort"
)

func testEdUpToPerm (l,s []int) bool {
	sort.Ints(l)
	sort.Ints(s) 
	if len(l) != len(s) {
		return false
	}
	for i := 0; i < len(l); i++ {
		if l[i] != s[i] {
			return false
		}
	}
	return true
}

func Test1(t *testing.T) {
	got := fastPower(5,2,2) 
	want := 4

	if got != want {
		t.Errorf("You said 2^2 = %d mod 5, but the answer is %d", got, want)
	}
}

func Test2(t *testing.T) {
	got := fastPower(15485863,2,15485862) 
	want := 1

	if got != want {
		t.Errorf("You said 2^(15485862) = %d mod 15485863, but the answer is %d", got, want)
	}
}

func Test3(t *testing.T) {
	got := primRootsModN(7)
	want := []int{3,5} 

	if !testEdUpToPerm(got,want) {
		t.Errorf("You said the primitive roots modulo 7 are %d, but the answer is %d", got, want)
	}
}

func Test4(t *testing.T) {
	got := primRootsModN(29)
	want := []int{2, 3, 8, 10, 11, 14, 15, 18, 19, 21, 26, 27} 

	if !testEdUpToPerm(got,want) {
		t.Errorf("You said the primitive roots modulo 29 are %d, but the answer is %d", got, want)
	}
}


