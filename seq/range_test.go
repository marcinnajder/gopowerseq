package seq_test

import (
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, slices.Collect(seq.Range(1, 3)))

	// test below does not work :| because of following implementation 'func Collect[E any](seq iter.Seq[E]) []E { return AppendSeq([]E(nil), seq }'
	// assert.Equal(t, []int{}, slices.Collect(seq.Range(1, 0)))

	s1 := slices.Collect(seq.Range(1, 0))
	assert.Equal(t, 0, len(s1))
}
