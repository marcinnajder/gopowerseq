package seq

import (
	"iter"
)

func GroupValueBy[T any, K comparable, V any](keyFunc Func[T, K], valueFunc Func[T, V]) OperatorR[T, map[K][]V] {
	return func(s iter.Seq[T]) map[K][]V {
		m := map[K][]V{}
		for v := range s {
			key := keyFunc(v)
			m[key] = append(m[key], valueFunc(v))
		}
		return m
	}
}
