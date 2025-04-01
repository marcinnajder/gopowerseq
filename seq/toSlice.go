package seq

import (
	"iter"
	"slices"
)

func ToSlice[T any](items iter.Seq[T]) []T {
	return slices.Collect(items)
}
