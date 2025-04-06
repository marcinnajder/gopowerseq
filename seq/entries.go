package seq

import (
	"iter"
	"maps"
)

func Entries[K comparable, V any](m map[K]V) iter.Seq[Tuple[K, V]] {
	return func(yield func(Tuple[K, V]) bool) {
		for key, value := range maps.All(m) {
			if !yield(Tuple[K, V]{key, value}) {
				return
			}
		}
	}
}
