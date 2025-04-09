package seq_test

import (
	"fmt"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

// assert.deepEqual([...pairwise([])], []);
// assert.deepEqual([...pairwise([1])], []);
// assert.deepEqual([...pairwise([1, 2])], [[1, 2]]);
// assert.deepEqual([...pairwise([1, 2, 3])], [[1, 2], [2, 3]]);

// assert.deepEqual([...pipe([1, 2, 3], pairwise(), map(([prev, next]) => `${prev}-${next}`))], ["1-2", "2-3"]);

func TestPairwise(t *testing.T) {
	Assert_EqualSeq2(t, []int{}, []int{}, seq.Pairwise[int]()(seq.Of[int]()))

	Assert_EqualSeq2(t, []int{}, []int{}, seq.Pairwise[int]()(seq.Of(1)))
	Assert_EqualSeq2(t, []int{1}, []int{2}, seq.Pairwise[int]()(seq.Of(1, 2)))
	Assert_EqualSeq2(t, []int{1, 2}, []int{2, 3}, seq.Pairwise[int]()(seq.Of(1, 2, 3)))

	pairs := seq.Pipe(
		seq.Of(1, 2, 3),
		seq.PairwiseFunc(func(prev, next int) string { return fmt.Sprint(prev, "-", next) }))

	Assert_EqualSeq(t, seq.Of("1-2", "2-3"), pairs)
}
