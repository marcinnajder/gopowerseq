package seq

import (
	"iter"
)

func Reduce[T any](f Func2[T, T, T]) OperatorR[T, T] {
	return func(s iter.Seq[T]) T {
		next, stop := iter.Pull(s)
		defer stop()

		result, hasVal := next()

		if !hasVal {
			panic("seq.Reduce: sequence contains no element")
		}

		for {
			val, hasVal2 := next()

			if !hasVal2 {
				break
			}

			result = f(result, val)
		}

		return result
	}
}

func ReduceA[T, A any](f Func2[A, T, A], initValue A) OperatorR[T, A] {
	return func(s iter.Seq[T]) A {
		result := initValue
		for val := range s {
			result = f(result, val)
		}
		return result
	}
}
