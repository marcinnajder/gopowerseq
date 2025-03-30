package seq

import (
	"iter"
)

func Take[T any](count int) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			if count <= 0 {
				return
			}
			for v := range s {
				if !yield(v) {
					return
				}
				count = count - 1
				if count <= 0 {
					return
				}
			}
		}
	}
}
