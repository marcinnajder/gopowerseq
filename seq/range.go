package seq

import (
	"iter"
)

func Range(start, count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		end := start + count
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
