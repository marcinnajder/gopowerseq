package seq_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

type Item struct {
	name  string
	items []int
}

func TestFlatMap(t *testing.T) {
	items := slices.Values([]Item{{"a", []int{}}, {"b", []int{1}}, {"c", []int{1, 2}}})

	getItems := func(item Item) iter.Seq[int] {
		return slices.Values(item.items)
	}

	formatResult := func(item Item, subitem int) string {
		return fmt.Sprint(item.name, subitem)
	}

	assert.Equal(t, []int{1, 1, 2}, slices.Collect(seq.FlatMap(getItems)(items)))
	assert.Equal(t, []string{"b1", "c1", "c2"}, slices.Collect(seq.FlatMapR(getItems, formatResult)(items)))
}
