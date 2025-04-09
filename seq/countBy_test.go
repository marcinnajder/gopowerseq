package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestCountBy(t *testing.T) {
	items := seq.Of("a", "a", "cc", "ddd")
	m := seq.CountBy(func(s string) int { return len(s) })(items)
	assert.Equal(t, 2, m[1])
	assert.Equal(t, 1, m[2])
	assert.Equal(t, 1, m[3])
}
