package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestFilter(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(2, 4), seq.Filter(IsEven)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Empty[int](), seq.Filter(IsEven)(seq.Of(1, 21, 3, 41)))
}
