package seq

import (
	"iter"
)

func TakeWhile[T any](f Func[T, bool]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			for item := range s {
				if f(item) {
					if !yield(item) {
						return
					}
				} else {
					break
				}
			}
		}
	}
}
