package seq

import "iter"

func Expand[T any](f func(T) iter.Seq[T]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			currentSeqs := []iter.Seq[T]{s}
			for {
				nextSeqs := make([]iter.Seq[T], 0)

				for _, currentSeq := range currentSeqs {
					for item := range currentSeq {
						if !yield(item) {
							return
						}

						if nextSeq := f(item); nextSeq != nil {
							nextSeqs = append(nextSeqs, nextSeq)
						}
					}
				}

				if len(nextSeqs) == 0 {
					return
				}

				currentSeqs = nextSeqs
			}
		}
	}
}
