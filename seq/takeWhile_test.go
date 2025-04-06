package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestTakeWhile(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2), seq.TakeWhile(func(i int) bool { return i < 3 })(seq.Of(1, 2)))
	Assert_EqualSeq(t, seq.Of(1), seq.TakeWhile(func(i int) bool { return i < 2 })(seq.Of(1, 2)))
	Assert_EqualSeq(t, seq.Of[int](), seq.TakeWhile(func(i int) bool { return i < 1 })(seq.Of(1, 2)))
}
