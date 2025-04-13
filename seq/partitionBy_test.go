package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestPartitionBy(t *testing.T) {
	Assert_EqualSeq(t, seq.Empty[[]int](), seq.PartitionBy(IsEven)(seq.Of[int]()))
	Assert_EqualSeq(t, seq.Of([]int{1}), seq.PartitionBy(IsEven)(seq.Of(1)))
	Assert_EqualSeq(t, seq.Of([]int{1}, []int{2}, []int{3}, []int{4}), seq.PartitionBy(IsEven)(seq.Of(1, 2, 3, 4)))
	Assert_EqualSeq(t, seq.Of([]int{1}, []int{2, 4, 6}, []int{3}, []int{4}), seq.PartitionBy(IsEven)(seq.Of(1, 2, 4, 6, 3, 4)))
}
