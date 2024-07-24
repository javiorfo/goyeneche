package functional

import "github.com/javiorfo/goyeneche/internal/interfaces"

func Filter[T comparable](arr []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T, R any](arr []T, mapper func(T) R) []R {
	result := make([]R, 0)
	for _, v := range arr {
		result = append(result, mapper(v))
	}
	return result
}

func Sum[T interfaces.TypeAdder](arr []T) T {
    var sum T
    for _, v := range arr {
       sum += v
    }
    return sum
}
