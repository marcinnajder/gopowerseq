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
