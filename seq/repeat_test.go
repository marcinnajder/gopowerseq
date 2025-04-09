package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestRepeat(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 1, 2, 3), seq.Repeat[int](2)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.Repeat[int](1)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of[int](), seq.Repeat[int](0)(seq.Of(1, 2, 3)))

	items := seq.Pipe2(
		seq.Of(1, 2, 3),
		seq.Repeat[int](-1),
		seq.Take[int](7))

	Assert_EqualSeq(t, seq.Of(1, 2, 3, 1, 2, 3, 1), items)

}
