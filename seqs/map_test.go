package seqs_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	assert.Equal(t, []int{2, 3, 4}, seqs.Map([]int{1, 2, 3}, Inc))
	assert.Equal(t, []string{"1", "2", "3"}, seqs.Map([]int{1, 2, 3}, ToString))
	// assert_EqualSeq(t, seq.Empty[int](), seq.Map(inc)(seq.Empty[int]()))
	// assert_EqualSeq(t, seq.Of("1", "2"), seq.Map(toString)(seq.Of(1, 2)))
}
