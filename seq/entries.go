package seq

import (
	"iter"
	"maps"
)

func Entries[K comparable, V any](m map[K]V) iter.Seq[Tuple2[K, V]] {
	return func(yield func(Tuple2[K, V]) bool) {
		for key, value := range maps.All(m) {
			if !yield(Tuple2[K, V]{key, value}) {
				return
			}
		}
	}
}
