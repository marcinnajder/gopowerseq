package seq

import (
	"iter"
)

func FlatMapR[T, C, R any](subcollFunc Func[T, iter.Seq[C]], resultFunc Func2[T, C, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for item := range s {
				for subitem := range subcollFunc(item) {
					if !yield(resultFunc(item, subitem)) {
						return
					}
				}
			}
		}
	}
}

func FlatMap[T, C any](subcollFunc Func[T, iter.Seq[C]]) Operator[T, C] {
	return FlatMapR(subcollFunc, func(item T, subitem C) C { return subitem })
}
