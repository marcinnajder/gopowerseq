package seq

import "iter"

func Memoize[T any]() Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		next, stop := iter.Pull(s)
		iters := 0

		results := make([]struct {
			value    T
			hasValue bool
		}, 0)

		return func(yield func(T) bool) {
			iters++

			defer func() {
				iters--
				if iters == 0 {
					iters = 1 // prevent from stopping twice
					stop()
				}
			}()

			for i := 0; ; i++ {
				if i < len(results) {
					result := results[i]
					if !result.hasValue {
						return
					}
					if !yield(result.value) {
						return
					}

					continue
				}

				if len(results) > 0 && !results[i-1].hasValue {
					return
				}

				value, hasValue := next()

				result := struct {
					value    T
					hasValue bool
				}{value, hasValue}

				results = append(results, result)
				if !result.hasValue {
					return
				}
				if !yield(result.value) {
					return
				}
			}
		}
	}
}
