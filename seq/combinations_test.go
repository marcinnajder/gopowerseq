package seq_test

// package seq_test // 'pullToPush' function is "private"

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

// assert.deepEqual([...combinations([1, 2, 3], -1)], []);
//     assert.deepEqual([...combinations([1, 2, 3], 5)], []);
//     assert.deepEqual([...combinations([1, 2, 3], 3)], [[1, 2, 3]]);

//     assert.deepEqual([...combinations([1, 2, 3], 1)], [[1], [2], [3]]);

//     assert.deepEqual([...combinations([1, 2, 3, 4], 2)], [[1, 2], [1, 3], [2, 3], [1, 4], [2, 4], [3, 4]]);

//     assert.deepEqual([...combinations([1, 2, 3, 4], 3)], [[1, 2, 3], [1, 2, 4], [1, 3, 4], [2, 3, 4]]);

//     assert.deepEqual([...combinations(range(1, 3), 3)], [[1, 2, 3]]);

func TestCombinations(t *testing.T) {
	Assert_EqualSeq(t, seq.Of([]int{1, 2, 3}), seq.Combinations[int](3)(seq.Of(1, 2, 3)))
	Assert_EqualSeq(t, seq.Of([]int{1}, []int{2}, []int{3}), seq.Combinations[int](1)(seq.Of(1, 2, 3)))

	Assert_EqualSeq(t,
		seq.Of([]int{1, 2}, []int{1, 3}, []int{2, 3}, []int{1, 4}, []int{2, 4}, []int{3, 4}),
		seq.Combinations[int](2)(seq.Of(1, 2, 3, 4)))

	Assert_EqualSeq(t,
		seq.Of([]int{1, 2, 3}, []int{1, 2, 4}, []int{1, 3, 4}, []int{2, 3, 4}),
		seq.Combinations[int](3)(seq.Of(1, 2, 3, 4)))
}
