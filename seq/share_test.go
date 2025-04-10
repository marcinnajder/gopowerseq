package seq_test

import (
	"iter"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestShare(t *testing.T) {
	items := seq.Share[int]()(seq.Of(1, 2, 3, 4))

	next1, _ := iter.Pull(items)
	next2, _ := iter.Pull(items)

	move := func(expectedValue int, expectedHasValue bool, next func() (int, bool)) {
		value, hasValue := next()
		assert.Equal(t, expectedHasValue, hasValue)
		assert.Equal(t, expectedValue, value)
	}

	move(1, true, next1)
	move(2, true, next1)
	move(3, true, next2)
	move(4, true, next1)
	move(0, false, next1)
	move(0, false, next2)
}
