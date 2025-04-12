package seq_test

import (
	"iter"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestMemoize(t *testing.T) {
	val := 0
	var next3Seq iter.Seq[int] = func(yield func(int) bool) {
		for i := 0; i < 3; i++ {
			val++
			if !yield(val) {
				return
			}
		}
	}

	Assert_EqualSeq(t, seq.Of(1, 2, 3), next3Seq)
	Assert_EqualSeq(t, seq.Of(4, 5, 6), next3Seq)

	next3SeqMemoized := seq.Memoize[int]()(next3Seq)

	Assert_EqualSeq(t, seq.Of(7, 8, 9), next3SeqMemoized)
	Assert_EqualSeq(t, seq.Of(7, 8, 9), next3SeqMemoized)
}
