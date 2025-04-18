package seq

import (
	"iter"
)

func Share[T any]() Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		next, _ := iter.Pull(s)
		//iters := 0

		return func(yield func(T) bool) {
			// - there is not ideal solution for cleaning resources :/
			// - finally cleaning logic has been commented out to be compatible with JS powerseq library

			// iters++
			// defer func() {
			// 	iters--
			// 	if iters == 0 {
			// 		stop()
			// 	}
			// }()

			for {
				value, hasValue := next()
				if !hasValue || !yield(value) {
					return
				}
			}
		}
	}
}
