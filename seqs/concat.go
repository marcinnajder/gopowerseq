package seqs

import (
	"iter"
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func Concat[T any](ss ...[]T) iter.Seq[T] {
	sss := Map(ss, slices.Values)
	return seq.Concat(sss...)
}
