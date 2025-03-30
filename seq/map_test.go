package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
)

func TestMap(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(2, 3), seq.Map(Inc)(seq.Of(1, 2)))
	Assert_EqualSeq(t, seq.Empty[int](), seq.Map(Inc)(seq.Empty[int]()))
	Assert_EqualSeq(t, seq.Of("1", "2"), seq.Map(ToString)(seq.Of(1, 2)))
}
