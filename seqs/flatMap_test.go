package seqs_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	entries := seq.Entries(map[string][]int{"a": {}, "b": {1}, "c": {1, 2}})

	getItems := func(e seq.Tuple[string, []int]) iter.Seq[int] {
		return slices.Values(e.Item2)
	}

	i := 0
	for item, subitem := range seqs.FlatMap(entries, getItems) {
		switch {
		case i == 0:
			assert.Equal(t, "b", item.Item1)
			assert.Equal(t, 1, subitem)
		case i == 1:
			assert.Equal(t, "c", item.Item1)
			assert.Equal(t, 1, subitem)
		case i == 2:
			assert.Equal(t, "c", item.Item1)
			assert.Equal(t, 2, subitem)
		}
		i++
	}
}
