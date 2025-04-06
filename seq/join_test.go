package seq_test

import (
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

type Item1 struct {
	id   int
	name string
}
type Item2 struct {
	id   int
	name string
}

func TestJoin(t *testing.T) {

	items1 := slices.Values([]Item1{{1, "one"}, {2, "two"}, {3, "three_"}, {3, "three__"}})
	items2 := slices.Values([]Item2{{1, "ONE"}, {3, "THREE"}, {4, "FOUR"}})

	getId1 := func(item Item1) int {
		return item.id
	}
	getId2 := func(item Item2) int {
		return item.id
	}

	i := 0
	for item1, item2 := range seq.Join(items1, items2, getId1, getId2) {
		switch {
		case i == 0:
			assert.Equal(t, "one", item1.name)
			assert.Equal(t, "ONE", item2.name)
		case i == 1:
			assert.Equal(t, "three_", item1.name)
			assert.Equal(t, "THREE", item2.name)
		case i == 2:
			assert.Equal(t, "three__", item1.name)
			assert.Equal(t, "THREE", item2.name)
		}
		i++
	}
}
