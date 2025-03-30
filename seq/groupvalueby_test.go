package seq_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestGroupValueBy(t *testing.T) {
	items := slices.Values([]string{"a", "aa", "b", "ccc", "bb"})

	m := seq.GroupValueBy(
		func(s string) int {
			return len(s)
		},
		func(s string) string {
			return fmt.Sprint(s, len(s))
		})(items)

	assert.Equal(t, map[int][]string{1: {"a1", "b1"}, 2: {"aa2", "bb2"}, 3: {"ccc3"}}, m)
}
