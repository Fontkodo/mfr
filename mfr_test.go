package mfr

import (
	"fmt"
	"strconv"
	"testing"
)

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
	stringXs := Map(strconv.Itoa, xs)
	if len(xs) != len(stringXs) {
		t.Error("Length of inputs and outputs not equal")
	}
	for i, x := range xs {
		if strconv.Itoa(x) != stringXs[i] {
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
	concatenation := Reduce(func(acc string, val int) string { return fmt.Sprintf("%s%d", acc, val) }, xs)
	if concatenation != "12345" {
		t.Error()
	}
	concatenation = Reduce(func(acc string, val int) string { return fmt.Sprintf("%s%d", acc, val) }, []int{})
	if concatenation != "" {
		t.Error()
	}
	concatenation = Reduce(func(acc string, val int) string { return fmt.Sprintf("%s%d", acc, val) }, xs[:1])
	if concatenation != "1" {
		t.Error()
	}
}

func compareSlices[ST comparable](s1 []ST, s2 []ST, t *testing.T) {
	if len(s1) != len(s2) {
		t.Error()
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			t.Error()
		}
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	keys := Keys(m)
	compareSlices[string]([]string{"a", "b"}, keys, t)
}
