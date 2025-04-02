package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.Equal(t, false, seq.All(IsEven)(seq.Of(2, 4, 3)))
	assert.Equal(t, true, seq.All(IsEven)(seq.Of(2, 4)))
	assert.Equal(t, true, seq.All(IsEven)(seq.Of[int]()))
}
