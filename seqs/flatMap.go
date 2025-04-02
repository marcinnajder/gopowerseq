package seqs

import (
	"iter"

	"github.com/marcinnajder/gopowerseq/seq"
)

func FlatMapR[T, C, R any](s iter.Seq[T], subcollFunc seq.Func[T, iter.Seq[C]], resultFunc seq.Func2[T, C, R]) iter.Seq[R] {
	return seq.FlatMapR(subcollFunc, resultFunc)(s)
}

func FlatMap[T, C any](s iter.Seq[T], subcollFunc seq.Func[T, iter.Seq[C]]) iter.Seq2[T, C] {
	return func(yield func(T, C) bool) {
		for item := range s {
			for subitem := range subcollFunc(item) {
				if !yield(item, subitem) {
					return
				}
			}
		}
	}
}
