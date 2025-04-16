package seq

import (
	"iter"
)

func SequenceEqual[T comparable](ss ...iter.Seq[T]) bool {
	if len(ss) < 2 {
		return true
	}

	nextFuncs, stopAll := getIterators(ss)
	defer stopAll()

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
