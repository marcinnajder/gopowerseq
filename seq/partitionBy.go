package seq

import (
	"iter"
)

func PartitionBy[T any, V comparable](f Func[T, V]) Operator[T, []T] {
	return func(s iter.Seq[T]) iter.Seq[[]T] {
		return func(yield func([]T) bool) {
			next, stop := iter.Pull(s)
			defer stop()

			value, hasValue := next()
			if !hasValue {
				return
			}

			pack := []T{value}
			prevKey := f(value)

			for {
				value, hasValue = next()
				if !hasValue {
					break
				}

				if key := f(value); key != prevKey {
					if !yield(pack) {
						return
					}

					pack = []T{value}
					prevKey = key
				} else {
					pack = append(pack, value)
				}
			}

			if len(pack) > 0 {
				if !yield(pack) {
					return
				}
			}
		}
	}
}
