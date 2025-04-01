package seq

import (
	"iter"
)

func FlatMapR[T, C, R any](collFunc Func[T, iter.Seq[C]], resultFunc Func2[T, C, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for item := range s {
				for subitem := range collFunc(item) {
					if !yield(resultFunc(item, subitem)) {
						return
					}
				}
			}
		}
	}
}

func FlatMap[T, C any](collFunc Func[T, iter.Seq[C]]) Operator[T, C] {
	return FlatMapR(collFunc, func(item T, subitem C) C { return subitem })
}
