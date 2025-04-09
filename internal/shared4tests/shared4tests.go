package shared4tests

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// it seems that tests for one packages can not import from tests from other package, thats way this package was introduces
// because this package is put into folder called "internal", it will not be avaliable from the outside of module 'gopowerseq'

func IsEven(i int) bool {
	return i%2 == 0
}

func Inc(i int) int {
	return i + 1
}

func Add(a, b int) int {
	return a + b
}

func ToString(i int) string {
	return strconv.Itoa(i)
}

func Assert_EqualSeq[T any](t *testing.T, expected iter.Seq[T], actual iter.Seq[T]) {
	assert.Equal(t, slices.Collect(expected), slices.Collect(actual))
}

func Assert_EqualSeq2[T1, T2 any](t *testing.T, e1 []T1, e2 []T2, actual iter.Seq2[T1, T2]) {
	items1, items2 := seq2To2Seq(actual)
	assert.Equal(t, e1, items1)
	assert.Equal(t, e2, items2)
}

func seq2To2Seq[T1, T2 any](s iter.Seq2[T1, T2]) ([]T1, []T2) {
	items1 := make([]T1, 0)
	items2 := make([]T2, 0)

	for item1, item2 := range s {
		items1 = append(items1, item1)
		items2 = append(items2, item2)
	}

	return items1, items2
}
