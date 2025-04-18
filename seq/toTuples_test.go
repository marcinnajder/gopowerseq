package seq_test

import (
	"maps"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestToTuples(t *testing.T) {
	m := map[int]string{1: "one"}
	Assert_EqualSeq(t,
		seq.Of(seq.Tuple[int, string]{1, "one"}),
		seq.Pipe(maps.All(m), seq.ToTuples[int, string]()))
}
