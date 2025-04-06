package seq

import (
	"iter"
)

func SkipWhile[T any](f Func[T, bool]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {

			next, stop := iter.Pull(s)
			defer stop()

			for {
				value, hasValue := next()

				if !hasValue {
					return
				}
				if !f(value) {
					if !yield(value) {
						return
					}
					break
				}
			}

			for {
				value, hasValue := next()
				if !hasValue || !yield(value) {
					return
				}
			}
		}
	}
}
