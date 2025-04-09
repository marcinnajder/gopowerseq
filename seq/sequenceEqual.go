package seq

import (
	"iter"
)

func SequenceEqual[T comparable](ss ...iter.Seq[T]) bool {
	if len(ss) < 2 {
		return true
	}

	nextFuncs := make([]func() (T, bool), len(ss))
	stopFuncs := make([]func(), len(ss))

	defer func() {
		for _, stop := range stopFuncs {
			stop()
		}
	}()

	for i, s := range ss {
		next, stop := iter.Pull(s)
		nextFuncs[i] = next
		stopFuncs[i] = stop
	}

	for {
		value1, hasValue1 := nextFuncs[0]()

		for i := 1; i < len(nextFuncs); i++ {
			value2, hasValue2 := nextFuncs[i]()
			if hasValue1 != hasValue2 || value1 != value2 {
				return false
			}
		}

		if !hasValue1 {
			break
		}
	}

	return true
}

func SequenceEqualP[T comparable](ss ...iter.Seq[T]) OperatorR[T, bool] {
	return func(s iter.Seq[T]) bool {
		sss := append([]iter.Seq[T]{s}, ss...)
		return SequenceEqual(sss...)
	}
}
