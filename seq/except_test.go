package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestExcept(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(2, 3), seq.Except(seq.Of(1, 2, 3, 4, 2, 4), seq.Of(4, 5, 6, 1)))

	Assert_EqualSeq(t,
		seq.Of(2),
		seq.ExceptFunc(seq.Of(1, 2, 3, 4, 2, 4), seq.Of(4, 6), func(i int) int { return i % 3 }))
}
