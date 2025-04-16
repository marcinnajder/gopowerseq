package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestInterleave(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.Interleave(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 10, 2, 20, 3, 30), seq.Interleave(seq.Of(1, 2, 3), seq.Of(10, 20, 30)))
	Assert_EqualSeq(t, seq.Of(1, 10, 2, 20, 3, 30), seq.Interleave(seq.Of(1, 2, 3, 4), seq.Of(10, 20, 30)))
	Assert_EqualSeq(t, seq.Of(1, 10, 100, 2, 20, 200), seq.Interleave(seq.Of(1, 2, 3), seq.Of(10, 20, 30), seq.Of(100, 200)))

	Assert_EqualSeq(t, seq.Of(1, 10, 2, 20, 3, 30), seq.InterleaveP(seq.Of(10, 20, 30))(seq.Of(1, 2, 3)))
}
