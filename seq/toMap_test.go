package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestToMap(t *testing.T) {
	m := seq.ToMap(func(item string) int { return len(item) })(seq.Of("a", "b", "cc", "ddd", "ee"))

	assert.Equal(t, "b", m[1])
	assert.Equal(t, "ee", m[2])
	assert.Equal(t, "ddd", m[3])
}
