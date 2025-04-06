package seq_test

import (
	"fmt"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	toString := func(n int, s string) string {
		return fmt.Sprintf("%d-%s", n, s)
	}
	res1 := seq.ZipFuncP(seq.Of("a", "b"), toString)(seq.Range(1, 10))
	Assert_EqualSeq(t, seq.Of("1-a", "2-b"), res1)

	toString3 := func(n int, s string, b bool) string {
		return fmt.Sprintf("%d-%s-%v", n, s, b)
	}

	res2 := seq.Zip3FuncP(seq.Of("a", "b"), seq.Of(true, false, false, true), toString3)(seq.Range(1, 10))
	Assert_EqualSeq(t, seq.Of("1-a-true", "2-b-false"), res2)

	//res3 := seq.Pipe(seq.Range(1, 10), seq.ZipP[int, string](seq.Of("a", "b")))

	res3 := seq.Zip(seq.Range(1, 10), seq.Of("a", "b"))
	for item1, item2 := range res3 {
		if item1 == 1 {
			assert.Equal(t, "a", item2)
		}
		if item1 == 2 {
			assert.Equal(t, "b", item2)
		}
	}

}
