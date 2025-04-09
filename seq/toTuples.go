package seq

import (
	"iter"
)

func ToTuples[T1, T2 any]() OperatorTR[iter.Seq2[T1, T2], iter.Seq[Tuple[T1, T2]]] {
	return func(s iter.Seq2[T1, T2]) iter.Seq[Tuple[T1, T2]] {
		return func(yield func(Tuple[T1, T2]) bool) {
			for item1, item2 := range s {
				if !yield(Tuple[T1, T2]{item1, item2}) {
					return
				}
			}
		}
	}
}
