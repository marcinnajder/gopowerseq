package seq

import (
	"iter"
	"slices"
)

func ToSlice[T any]() OperatorR[T, []T] {
	return func(s iter.Seq[T]) []T {
		return slices.Collect(s)
	}
}
