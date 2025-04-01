package seqs_test

import (
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 10, 20, 30}, slices.Collect(seqs.Concat([]int{1, 2, 3}, []int{10, 20, 30})))
}
