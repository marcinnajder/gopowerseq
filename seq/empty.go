package seq

import (
	"iter"
)

func emptySeq[T any](yield func(T) bool) {
}

func Empty[T any]() iter.Seq[T] {
	return emptySeq
}
