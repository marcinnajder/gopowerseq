package seq

import (
	"iter"
)

// func Share[T any]() Operator[T, T] {
// 	return func(s iter.Seq[T]) iter.Seq[T] {
// 		next, stop := iter.Pull(s)
// 		return pullToPush(next, stop)
// 	}
// }

func Share[T any]() Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		next, stop := iter.Pull(s)
		iters := 0

		return func(yield func(T) bool) {
			iters++

			defer func() {
				iters--
				if iters == 0 {
					stop()
				}
			}()

			for {
				value, hasValue := next()
				if !hasValue || !yield(value) {
					return
				}
			}
		}
	}
}
