package seq

import (
	"iter"

	"github.com/marcinnajder/gopowerseq/sequ"
)

func distinct[T any, K comparable, R any](f Func[T, K], resultFunc Func2[T, K, R]) Operator[T, R] {
	return func(s iter.Seq[T]) iter.Seq[R] {
		return func(yield func(R) bool) {
			set := make(map[K]struct{})
			for item := range s {
				key := f(item)
				if _, ok := set[key]; !ok {
					set[key] = struct{}{}
					if !yield(resultFunc(item, key)) {
						return
					}
				}
			}
		}
	}
}

func DistinctFunc[T any, K comparable](f Func[T, K]) Operator[T, K] {
	return distinct(f, func(item T, key K) K { return key })
}

func Distinct[T comparable]() Operator[T, T] {
	return DistinctFunc(sequ.Identity[T])
}

func DistinctBy[T any, K comparable](f Func[T, K]) Operator[T, T] {
	return distinct(f, func(item T, key K) T { return item })
}

func DistinctUntilChangedFunc[T any, K comparable](f Func[T, K]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return func(yield func(T) bool) {
			next, stop := iter.Pull(s)
			defer stop()

			value, hasValue := next()
			if !hasValue {
				return
			}

			lastKey := f(value)
			if !yield(value) {
				return
			}

			for {
				value, hasValue := next()
				if !hasValue {
					return
				}
				key := f(value)
				if lastKey != key {
					lastKey = key
					if !yield(value) {
						return
					}
				}
			}
		}
	}
}

func DistinctUntilChanged[T comparable]() Operator[T, T] {
	return DistinctUntilChangedFunc(sequ.Identity[T])
}
