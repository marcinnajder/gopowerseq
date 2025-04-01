package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	val1, index1 := seq.Find(IsEven)(seq.Of(1, 2, 3, 4))
	assert.Equal(t, 1, index1)
	assert.Equal(t, 2, val1)

	val2, index2 := seq.Find(IsEven)(seq.Of(1, 21, 3, 41))
	assert.Equal(t, -1, index2)
	assert.Equal(t, 0, val2)
}
