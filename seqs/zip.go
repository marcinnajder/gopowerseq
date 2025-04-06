package seqs

import (
	"iter"
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func Zip[T1, T2 any](s1 []T1, s2 []T2) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		pairs := seq.ZipP(slices.Values(s2), func(item1 T1, item2 T2) seq.Tuple[T1, T2] {
			return seq.Tuple[T1, T2]{Item1: item1, Item2: item2}
		})(slices.Values(s1))

		for pair := range pairs {
			if !yield(pair.Item1, pair.Item2) {
				return
			}
		}
	}
}
