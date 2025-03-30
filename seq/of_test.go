package seq_test

import (
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, slices.Collect(seq.Of(1, 2, 3)))

	// it is strange but this test will not pass, detail explanation is in 'range_test.go' file
	// assert.Equal(t, []int{}, slices.Collect(seq.Of[int]())) ->

	s1 := slices.Collect(seq.Of[int]())
	assert.Equal(t, 0, len(s1))
}
