package seq

import (
	"iter"
)

func Filter[T any](f Func[T, bool]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			for v := range s {
				if f(v) {
					if !yield(v) {
						return
					}
				}
			}
		}
	}
}
