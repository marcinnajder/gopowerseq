package seq_test

// package seq_test // 'pullToPush' function is "private"

import (
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestPipe(t *testing.T) {
	assert.Equal(t, 2, seq.Pipe(1, Inc))
	assert.Equal(t, 3, seq.Pipe2(1, Inc, Inc))
	assert.Equal(t, "3", seq.Pipe3(1, Inc, Inc, ToString))
	assert.Equal(t, true, seq.Pipe4(1, Inc, Inc, Inc, IsEven))
}
