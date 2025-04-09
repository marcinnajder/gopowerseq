package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

func TestWindowed(t *testing.T) {
	Assert_EqualSeq(t, seq.Of([]int{5, 10, 15}, []int{10, 15, 100}), seq.Windowed[int](3)(seq.Of(5, 10, 15, 100)))
	Assert_EqualSeq(t, seq.Of([]int{5, 10, 15}), seq.Windowed[int](3)(seq.Of(5, 10, 15)))
	Assert_EqualSeq(t, seq.Of([]int{5}, []int{10}, []int{15}), seq.Windowed[int](1)(seq.Of(5, 10, 15)))

	Assert_EqualSeq(t, seq.Of[[]int](), seq.Windowed[int](3)(seq.Of(5, 10)))
	Assert_EqualSeq(t, seq.Of[[]int](), seq.Windowed[int](3)(seq.Of(5)))
	Assert_EqualSeq(t, seq.Of[[]int](), seq.Windowed[int](3)(seq.Of[int]()))
}
