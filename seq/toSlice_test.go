package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestToSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, seq.ToSlice(seq.Of(1, 2, 3)))

	s := seq.ToSlice(seq.Of[int]())
	assert.Equal(t, 0, len(s))

}
