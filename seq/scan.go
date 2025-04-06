package seq

import (
	"iter"
)

func Scan[T, A any](f Func2[A, T, A], initValue A) Operator[T, A] {
	return func(s iter.Seq[T]) iter.Seq[A] {
		return func(yield func(A) bool) {
			next := initValue
			if !yield(next) {
				return
			}
			for val := range s {
				next = f(next, val)
				if !yield(next) {
					return
				}
			}
		}
	}
}
