package seq

import (
	"iter"
)

func Skip[T any](count int) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			if count < 0 {
				return
			}

			next, stop := iter.Pull(s)
			defer stop()

			i := 0
			for i < count {
				if _, hasValue := next(); !hasValue {
					return
				}
				i++
			}
			for {
				value, hasValue := next()
				if !hasValue {
					return
				}
				if !yield(value) {
					return
				}
			}
		}
	}
}
