package seq

import "iter"

func GroupByV[T any, K comparable, V any](keyFunc Func[T, K], valueFunc Func[T, V]) OperatorR[T, map[K][]V] {
	return func(s iter.Seq[T]) map[K][]V {
		m := map[K][]V{}
		for v := range s {
			key := keyFunc(v)
			m[key] = append(m[key], valueFunc(v))
		}
		return m
	}
}

func GroupBy[T any, K comparable](f Func[T, K]) OperatorR[T, map[K][]T] {
	return GroupByV(f, identity)
}
