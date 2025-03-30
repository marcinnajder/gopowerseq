package seq

import (
	"iter"
)

func Map[T, R any](f Func[T, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v := range s {
				if !yield(f(v)) {
					return
				}
			}
		}
	}
}
