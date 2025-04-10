package seq

import (
	"iter"
)

func RepeatValue[T any](value T, count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if count < 0 { // infinite
			for {
				if !yield(value) {
					return
				}
			}
		} else {
			for i := 0; i < count; i++ {
				if !yield(value) {
					return
				}
			}
		}
	}
}
