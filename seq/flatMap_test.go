package seq_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

// it('flatmap', function () {
//     const items = [
//         { name: "a", items: [] },
//         { name: "b", items: [1] },
//         { name: "c", items: [1, 2], },
//     ];

//     assert.deepEqual([...flatmap(items, x => x.items)], [1, 1, 2]);
//     assert.deepEqual([...flatmap(items, x => x.items)], [1, 1, 2]);
//     assert.deepEqual([...flatmap(items, (x, index) => [index].concat(x.items))], [0, 1, 1, 2, 1, 2]);
//     assert.deepEqual([...flatmap(items, x => x.items, (item, subitem) => item.name + subitem)], ["b1", "c1", "c2"]);

//     assert.deepEqual([...flatmap<typeof items[0], number>(x => x.items)(items)], [1, 1, 2]);
// });

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
