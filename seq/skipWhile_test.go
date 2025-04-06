package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestSkipWhile(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(3, 4), seq.SkipWhile(func(i int) bool { return i < 3 })(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4), seq.SkipWhile(func(i int) bool { return i < 1 })(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of[int](), seq.SkipWhile(func(i int) bool { return i < 5 })(seq.Of(1, 2, 3, 4)))
}
