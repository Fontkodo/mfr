// package mfr offers standard functional operations:
// Map, Filter, Reduce
package mfr

// Map inputs to outputs using a transformation.
//
// Example
//
//	xs := []int{1, 2, 3, 4, 5}
//	squares := Map(func(x int) int { return x * x }, xs)
func Map[T1 any, T2 any](f func(T1) T2, items []T1) []T2 {
	result := make([]T2, 0)
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

// Filter inputs using a predicate.
//
// Example
//
//	xs := []int{1, 2, 3, 4, 5}
//	even := Filter(func(x int) bool { return x%2 == 0 }, xs)
func Filter[T any](f func(T) bool, items []T) []T {
	result := make([]T, 0)
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

// Reduce inputs to a scalar.
//
// Example
//
//	xs := []int{1, 2, 3, 4, 5}
//	sum := Reduce(func(x int, y int) int { return x + y }, xs)
func Reduce[T1 any, T2 any](f func(T2, T1) T2, items []T1) T2 {
	var acc T2
	for _, item := range items {
		acc = f(acc, item)
	}
	return acc
}
