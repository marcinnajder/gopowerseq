package seq_test

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.Equal(t, 3, seq.Count[int]()(seq.Of(1, 2, 3)))
	assert.Equal(t, 0, seq.Count[int]()(seq.Of[int]()))
}

func TestCountFunc(t *testing.T) {
	assert.Equal(t, 2, seq.CountFunc(IsEven)(seq.Of(1, 2, 3, 4, 5)))
	assert.Equal(t, 0, seq.CountFunc(IsEven)(seq.Of(1, 3, 5)))
}
