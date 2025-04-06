package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestDistinct(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3, 4), seq.Distinct[int]()(seq.Of(1, 2, 3, 4, 2, 4)))
}

func TestDistinctFunc(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 0), seq.DistinctFunc(mod3)(seq.Of(1, 2, 3, 4, 5, 6)))
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.DistinctFunc(getLength)(seq.Of("a", "aa", "ab", "abc")))
}

func TestDistinctBy(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 3), seq.DistinctBy(mod3)(seq.Of(1, 2, 3, 4, 5, 6)))
	Assert_EqualSeq(t, seq.Of(6, 5, 4), seq.DistinctBy(mod3)(seq.Of(6, 5, 4, 3, 2, 1)))
	Assert_EqualSeq(t, seq.Of("a", "aa", "abc"), seq.DistinctBy(getLength)(seq.Of("a", "aa", "ab", "abc")))
}

func TestDistinctUntilChanged(t *testing.T) {
	Assert_EqualSeq(t, seq.Of(1, 2, 1, 2, 3, 4, 2), seq.DistinctUntilChanged[int]()(seq.Of(1, 1, 2, 2, 2, 1, 2, 3, 3, 4, 2, 2)))
	Assert_EqualSeq(t, seq.Of("a", "aa", "ccc", "d"), seq.DistinctUntilChangedFunc(getLength)(seq.Of("a", "aa", "ab", "ccc", "d", "e")))
}

func mod3(x int) int {
	return x % 3
}

func getLength(x string) int {
	return len(x)
}
