package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestEntries(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	Assert_EqualSeq(t, seq.Of(seq.Tuple[string, int]{"one", 1}, seq.Tuple[string, int]{"two", 2}), seq.Entries(m))
}
