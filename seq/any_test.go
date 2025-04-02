package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	assert.Equal(t, true, seq.Any(IsEven)(seq.Of(1, 3, 4)))
	assert.Equal(t, false, seq.Any(IsEven)(seq.Of(1, 3)))
	assert.Equal(t, false, seq.Any(IsEven)(seq.Of[int]()))
}
