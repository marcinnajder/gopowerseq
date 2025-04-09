package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {

	assert.Equal(t, 2, seq.Average[int]()(seq.Of(1, 2, 3)))
	assert.Equal(t, 1, seq.Average[int]()(seq.Of(1)))
	assert.Equal(t, 0, seq.Average[int]()(seq.Of[int]()))
	assert.Equal(t, 2.0, seq.Average[float64]()(seq.Of(1.5, 2.5, 2.0)))
	assert.Equal(t, 2, seq.AverageFunc(func(item string) int { return len(item) })(seq.Of("a", "aa", "aaa")))

	// Assert_EqualSeq(t, seq.Of([]int{5, 10, 15}, []int{100}), seq.Chunk[int](3)(seq.Of(5, 10, 15, 100)))
	// Assert_EqualSeq(t, seq.Of([]int{5, 10, 15, 100}), seq.Chunk[int](4)(seq.Of(5, 10, 15, 100)))
	// Assert_EqualSeq(t, seq.Of([]int{5, 10, 15, 100}), seq.Chunk[int](5)(seq.Of(5, 10, 15, 100)))

	// Assert_EqualSeq(t, seq.Of([]int{5}, []int{10}, []int{15}, []int{100}), seq.Chunk[int](1)(seq.Of(5, 10, 15, 100)))
}
