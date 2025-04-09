package seq

import (
	"iter"
)

func CountBy[T any, K comparable](f Func[T, K]) OperatorR[T, map[K]int] {
	return func(s iter.Seq[T]) map[K]int {
		result := make(map[K]int)
		for v := range s {
			key := f(v)
			result[key] = (result[key] + 1)

		}
		return result
	}
}
