package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 10, seq.Sum[int]()(seq.Of(1, 2, 3, 4)))
	assert.Equal(t, "1234", seq.Sum[string]()(seq.Of("1", "2", "3", "4")))
	assert.Equal(t, 7, seq.SumFunc(func(s string) int { return len(s) })(seq.Of("a", "aa", "bbb", "c")))
}
