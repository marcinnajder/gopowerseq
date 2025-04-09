package seq

import (
	"iter"
)

func Repeat[T any](count int) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			if count < 0 { // infinite
				for {
					for v := range s {
						if !yield(v) {
							return
						}
					}
				}
			} else {
				for i := 0; i < count; i++ {
					for v := range s {
						if !yield(v) {
							return
						}
					}
				}
			}
		}
	}
}
