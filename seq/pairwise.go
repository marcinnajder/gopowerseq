package seq

import (
	"iter"
)

func Pairwise[T any]() OperatorR[T, iter.Seq2[T, T]] {
	return func(s iter.Seq[T]) iter.Seq2[T, T] {
		return func(yield func(T, T) bool) {
			next, stop := iter.Pull(s)
			defer stop()

			prev, hasPrev := next()
			if !hasPrev {
				return
			}

			for {
				current, hasCurrent := next()
				if !hasCurrent {
					return
				}

				if !yield(prev, current) {
					return
				}
				prev = current
			}
		}
	}
}

func PairwiseFunc[T, R any](f Func2[T, T, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for item1, item2 := range Pairwise[T]()(s) {
				if !yield(f(item1, item2)) {
					return
				}
			}
		}
	}
}
