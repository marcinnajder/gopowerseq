package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestConcat(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4, 5, 6), seq.Concat(seq.Of(1, 2, 3), seq.Of(4), seq.Of[int](), seq.Of(5, 6)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4, 5, 6), seq.ConcatP(seq.Of(4), seq.Of[int](), seq.Of(5, 6))(seq.Of(1, 2, 3)))
}
