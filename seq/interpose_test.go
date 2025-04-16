package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestInterpose(t *testing.T) {
	Assert_EqualSeq(t, seq.Of[int](), seq.Interpose(0)(seq.Of[int]()))
	Assert_EqualSeq(t, seq.Of(1), seq.Interpose(0)(seq.Of(1)))
	Assert_EqualSeq(t, seq.Of(1, 0, 2), seq.Interpose(0)(seq.Of(1, 2)))
	Assert_EqualSeq(t, seq.Of(1, 0, 2, 0, 3), seq.Interpose(0)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 0, 2, 0, 3), seq.Interpose(0)(seq.Of(1, 2, 3)))
}
