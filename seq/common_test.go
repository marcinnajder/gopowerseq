package seq

// package seq_test // 'pullToPush' function is "private"

import (
	"iter"
	"slices"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	assert.Equal(t, 2, Pipe(1, Inc))
	assert.Equal(t, 3, Pipe2(1, Inc, Inc))
	assert.Equal(t, "3", Pipe3(1, Inc, Inc, ToString))
	assert.Equal(t, true, Pipe4(1, Inc, Inc, Inc, IsEven))
}

func TestPullToPush(t *testing.T) {
	items := slices.Values([]int{1, 2, 3, 4})
	next, stop := iter.Pull(items)
	Assert_EqualSeq(t, Of(1, 2, 3, 4), pullToPush(next, stop))
}
