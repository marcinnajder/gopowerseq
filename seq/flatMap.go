package seq

import (
	"iter"
)

func FlatMapR[T, C, R any](subcoll Func[T, iter.Seq[C]], result Func2[T, C, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for item := range s {
				for subitem := range subcoll(item) {
					if !yield(result(item, subitem)) {
						return
					}
				}
			}
		}
	}
}

func FlatMap[T, C any](subcoll Func[T, iter.Seq[C]]) Operator[T, C] {
	return FlatMapR(subcoll, func(item T, subitem C) C { return subitem })
}
