package seq

import (
	"iter"
)

// export function interleave<T>(...iterables: Iterable<T>[]): Iterable<T> {
//     return flatmap(zip(...iterables, (...results: any[]) => results as Iterable<T>) as Iterable<Iterable<T>>, identity);
// }

func Interleave[T any](ss ...iter.Seq[T]) iter.Seq[T] {

	return func(yield func(T) bool) {
		nextFuncs, stopAll := getIterators(ss)
		defer stopAll()

		values := make([]T, len(ss))
		for {
			for i, nextFunc := range nextFuncs {
				value, hasValue := nextFunc()

				if !hasValue {
					return
				}

				values[i] = value
			}

			for _, value := range values {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func InterleaveP[T any](ss ...iter.Seq[T]) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		sss := append([]iter.Seq[T]{s}, ss...)
		return Interleave(sss...)
	}
}
