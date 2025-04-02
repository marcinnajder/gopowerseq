package seqs_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	i := 0
	for item1, item2 := range seqs.Zip([]string{"a", "b"}, []int{1, 2}) {
		switch {
		case i == 0:
			assert.Equal(t, "a", item1)
			assert.Equal(t, 1, item2)
		case i == 1:
			assert.Equal(t, "b", item1)
			assert.Equal(t, 2, item2)
		}
		i++
	}
}
