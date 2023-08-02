package mfr

import "testing"

func TestMap(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	squares := Map(func(x int) int { return x * x }, xs)
	if len(xs) != len(squares) {
		t.Error("Length of inputs and outputs not equal")
	}
	for i, x := range xs {
		if x*x != squares[i] {
			t.Error()
		}
	}
}

func TestFilter(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	even := Filter(func(x int) bool { return x%2 == 0 }, xs)
	if len(xs) < len(even) {
		t.Error("Length of subset greater than inputs")
	}
}

func TestReduce(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	sum := Reduce(func(x int, y int) int { return x + y }, xs)
	if sum != 15 {
		t.Error()
	}
	sum = Reduce(func(x int, y int) int { return x + y }, []int{})
	if sum != 0 {
		t.Error()
	}
	sum = Reduce(func(x int, y int) int { return x + y }, xs[:1])
	if sum != 1 {
		t.Error()
	}
}
