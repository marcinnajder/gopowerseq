package seq_test

import (
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	items := slices.Values([]string{"a", "aa", "b", "ccc", "bb"})
	m := seq.GroupBy(func(s string) int {
		return len(s)
	})(items)
	assert.Equal(t, map[int][]string{1: {"a", "b"}, 2: {"aa", "bb"}, 3: {"ccc"}}, m)
}
