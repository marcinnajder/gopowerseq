package seq

import (
	"iter"
	"slices"
)

func Of[T any](items ...T) iter.Seq[T] {
	return slices.Values(items)
}
