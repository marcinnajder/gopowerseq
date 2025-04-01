package seqs_test

import (
	"fmt"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {

	for item1, item2 := range seqs.Zip(seq.Of("a", "b"), seq.Range(1, 10)) {
		fmt.Println(item1, item2)
	}

	assert.Equal(t, 1, 2)
	// assert.Equal(t, []string{"1", "2", "3"}, seqs.Map([]int{1, 2, 3}, ToString))
	// assert_EqualSeq(t, seq.Empty[int](), seq.Map(inc)(seq.Empty[int]()))
	// assert_EqualSeq(t, seq.Of("1", "2"), seq.Map(toString)(seq.Of(1, 2)))
}
