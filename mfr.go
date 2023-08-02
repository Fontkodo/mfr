package mfr

func Map[T1 any, T2 any](f func(T1) T2, items []T1) []T2 {
	result := make([]T2, 0)
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

func Filter[T any](f func(T) bool, items []T) []T {
	result := make([]T, 0)
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}

func Reduce[T1 any, T2 any](f func(T2, T1) T2, items []T1) T2 {
	var acc T2
	for _, item := range items {
		acc = f(acc, item)
	}
	return acc
}
