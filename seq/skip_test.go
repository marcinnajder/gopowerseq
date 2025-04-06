package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestSkip(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(2, 3, 4), seq.Skip[int](1)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of(4), seq.Skip[int](3)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of[int](), seq.Skip[int](4)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of[int](), seq.Skip[int](5)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4), seq.Skip[int](0)(seq.Of(1, 2, 3, 4)))
}
