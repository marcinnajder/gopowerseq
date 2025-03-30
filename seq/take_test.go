package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
)

func TestTake(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2), seq.Take[int](2)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.Take[int](3)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.Take[int](4)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of[int](), seq.Take[int](0)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 2), seq.Pipe(seq.Of(1, 2, 3), seq.Take[int](2)))
}
