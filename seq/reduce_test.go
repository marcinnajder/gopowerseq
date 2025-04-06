package seq_test

import (
	"fmt"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, 30, seq.Reduce(Add)(seq.Of(5, 10, 15)))
	assert.Equal(t, 5, seq.Reduce(Add)(seq.Of(5)))
	assert.Panics(t, func() {
		seq.Reduce(Add)(seq.Of[int]())
	})
}

func TestReduceA(t *testing.T) {
	assert.Equal(t, 30, seq.ReduceA(Add, 0)(seq.Of(5, 10, 15)))
	assert.Equal(t, 5, seq.ReduceA(Add, 0)(seq.Of(5)))
	assert.Equal(t, 0, seq.ReduceA(Add, 0)(seq.Of[int]()))

	assert.Equal(t, "0-5-10-15", seq.ReduceA(func(prev string, current int) string {
		return fmt.Sprintf("%s-%d", prev, current)
	}, "0")(seq.Of(5, 10, 15)))
}
