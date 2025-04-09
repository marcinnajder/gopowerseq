package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestChunk(t *testing.T) {
	Assert_EqualSeq(t, seq.Of([]int{5, 10}, []int{15, 100}), seq.Chunk[int](2)(seq.Of(5, 10, 15, 100)))
	Assert_EqualSeq(t, seq.Of([]int{5, 10, 15}, []int{100}), seq.Chunk[int](3)(seq.Of(5, 10, 15, 100)))
	Assert_EqualSeq(t, seq.Of([]int{5, 10, 15, 100}), seq.Chunk[int](4)(seq.Of(5, 10, 15, 100)))
	Assert_EqualSeq(t, seq.Of([]int{5, 10, 15, 100}), seq.Chunk[int](5)(seq.Of(5, 10, 15, 100)))

	Assert_EqualSeq(t, seq.Of([]int{5}, []int{10}, []int{15}, []int{100}), seq.Chunk[int](1)(seq.Of(5, 10, 15, 100)))
}
