package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestRepeatValue(t *testing.T) {
	Assert_EqualSeq(t, seq.Of("txt", "txt", "txt"), seq.RepeatValue("txt", 3))

	items := seq.Pipe(
		seq.RepeatValue("txt", -1),
		seq.Take[string](4))

	Assert_EqualSeq(t, seq.Of("txt", "txt", "txt", "txt"), items)

}
