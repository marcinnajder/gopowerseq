package seqs

import (
	"iter"

	"github.com/marcinnajder/gopowerseq/seq"
)

func Zip[T1, T2 any](s1 iter.Seq[T1], s2 iter.Seq[T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		pairs := seq.ZipP(s2, func(item1 T1, item2 T2) seq.Tuple2[T1, T2] {
			return seq.Tuple2[T1, T2]{Item1: item1, Item2: item2}
		})(s1)

		for pair := range pairs {
			if !yield(pair.Item1, pair.Item2) {
				return
			}
		}
	}
}
