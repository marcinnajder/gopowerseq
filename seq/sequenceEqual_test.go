package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestSequenceEqual(t *testing.T) {
	assert.Equal(t, true, seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2, 3), seq.Of(1, 2, 3)))
	assert.Equal(t, true, seq.SequenceEqual(seq.Of[int](), seq.Of[int](), seq.Of[int]()))

	assert.Equal(t, false, seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2), seq.Of(1, 2)))
	assert.Equal(t, false, seq.SequenceEqual(seq.Of(1, 2), seq.Of(1, 22)))
}
