package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestEmpty(t *testing.T) {
	Assert_EqualSeq(t, seq.Of[int](), seq.Empty[int]())
}
