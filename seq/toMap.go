package seq

import (
	"iter"
)

func ToMap[T any, K comparable](keyFunc Func[T, K]) OperatorR[T, map[K]T] {
	return func(s iter.Seq[T]) map[K]T {
		m := make(map[K]T)
		for item := range s {
			m[keyFunc(item)] = item
		}
		return m
	}
}

func ToMapV[T any, K comparable, V any](keyFunc Func[T, K], valueFunc Func[T, V]) OperatorR[T, map[K]V] {
	return func(s iter.Seq[T]) map[K]V {
		m := make(map[K]V)
		for item := range s {
			m[keyFunc(item)] = valueFunc(item)
		}
		return m
	}
}
