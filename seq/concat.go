package seq

import (
	"iter"
)

func Concat[T any](s iter.Seq[T]) func(...iter.Seq[T]) iter.Seq[T] {
	return func(ss ...iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			for v := range s {
				if !yield(v) {
					return
				}
			}
			for _, s := range ss {
				for v := range s {
					if !yield(v) {
						return
					}
				}
			}
		}
	}
}
