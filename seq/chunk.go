package seq

import (
	"iter"
)

func Chunk[T any](size int) OperatorR[T, iter.Seq[[]T]] {
	if size <= 0 {
		panic("seq.Chunk: 'size' must be positive")
	}

	return func(s iter.Seq[T]) iter.Seq[[]T] {
		return func(yield func([]T) bool) {
			next, stop := iter.Pull(s)
			defer stop()

			for {
				var array = make([]T, size)

				for i := 0; i < size; i++ {
					value, hasValue := next()

					if !hasValue {
						if i > 0 {
							if !yield(array[0:i]) {
								return
							}
						}
						return
					}

					array[i] = value
				}

				if !yield(array) {
					return
				}
			}
		}
	}
}
