package seq

import (
	"iter"
)

func Take[T any](count int) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			i := count
			if i <= 0 {
				return
			}
			for v := range s {
				if !yield(v) {
					return
				}
				i = i - 1
				if i <= 0 {
					return
				}
			}
		}
	}
}
