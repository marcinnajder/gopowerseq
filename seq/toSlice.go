package seq

import (
	"iter"
	"slices"
)

func ToSlice[T any](s iter.Seq[T]) []T {
	return slices.Collect(s)
}
