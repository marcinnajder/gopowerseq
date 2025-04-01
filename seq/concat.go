package seq

import (
	"iter"
)

func Concat[T any](ss ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, s := range ss {
			for v := range s {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func ConcatP[T any](ss ...iter.Seq[T]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		sss := append([]iter.Seq[T]{s}, ss...)
		return Concat(sss...)
	}
}
