package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestUnion(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4, 5, 6), seq.Union(seq.Of(1, 2, 3, 4, 2, 4), seq.Of(4, 5, 6)))
	Assert_EqualSeq(t,
		seq.Of(4, 2, 3),
		seq.UnionFunc(seq.Of(4, 1, 1, 2, 4), seq.Of(1, 2, 3, 4, 5, 6), func(i int) int { return i % 3 }))
}
