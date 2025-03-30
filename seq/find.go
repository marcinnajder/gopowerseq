package seq

import (
	"iter"
)

func Find[T any](f Func[T, bool]) func(iter.Seq[T]) (val T, index int) {
	return func(s iter.Seq[T]) (T, int) {
		i := 0
		for v := range s {
			if f(v) {
				return v, i
			}
			i++
		}
		var zero T
		return zero, -1
	}
}
