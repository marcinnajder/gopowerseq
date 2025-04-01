package seq_test

import (
	"fmt"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestZip(t *testing.T) {
	toString := func(n int, s string) string {
		return fmt.Sprintf("%d-%s", n, s)
	}
	res1 := seq.ZipP(seq.Of("a", "b"), toString)(seq.Range(1, 10))
	Assert_EqualSeq(t, seq.Of("1-a", "2-b"), res1)

	toString3 := func(n int, s string, b bool) string {
		return fmt.Sprintf("%d-%s-%v", n, s, b)
	}

	res2 := seq.Zip3P(seq.Of("a", "b"), seq.Of(true, false, false, true), toString3)(seq.Range(1, 10))
	Assert_EqualSeq(t, seq.Of("1-a-true", "2-b-false"), res2)

}
