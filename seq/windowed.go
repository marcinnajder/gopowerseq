package seq

import (
	"iter"
	"slices"
)

// https://github.com/marcinnajder/misc/blob/master/2022_04_23_advent_of_code_in_csharp/AdventOfCode/Common/Extensions.cs
// Chunk LINQ ?

func Windowed[T any](size int) OperatorR[T, iter.Seq[[]T]] {
	if size <= 0 {
		panic("seq.Windowed: 'size' must be positive")
	}

	return func(s iter.Seq[T]) iter.Seq[[]T] {
		return func(yield func([]T) bool) {
			next, stop := iter.Pull(s)
			defer stop()

			window := make([]T, size)

			var i int
			for i = 0; i < size; i++ {
				value, hasValue := next()
				if !hasValue {
					break
				}
				window[i] = value
			}

			if i != size {
				return
			}

			if !yield(slices.Clone(window)) {
				return
			}

			mod := 0
			for {
				value, hasValue := next()
				if !hasValue {
					break
				}

				window[mod] = value
				mod = (mod + 1) % size

				var array = make([]T, size)
				for i = 0; i < size; i++ {
					array[i] = window[(mod+i)%size]
				}

				if !yield(array) {
					return
				}
			}
		}
	}
}
