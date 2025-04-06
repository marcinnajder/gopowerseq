package seq_test

import (
	"fmt"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestScan(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(0, 5, 15, 30), seq.Scan(Add, 0)(seq.Of(5, 10, 15)))
	Assert_EqualSeq(t, seq.Of(0), seq.Scan(Add, 0)(seq.Of[int]()))

	Assert_EqualSeq(t, seq.Of("0, 5, 15, 30"), seq.Scan(func(prev string, current int) string {
		return fmt.Sprintf("%s-%d", prev, current)
	}, "0")(seq.Of(5, 10, 15)))
}
