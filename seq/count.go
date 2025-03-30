package seq

import (
	"iter"
)

func Count[T any]() OperatorR[T, int] {
	return func(s iter.Seq[T]) int {
		count := 0
		for range s {
			count++
		}
		return count
	}
}

func CountFunc[T any](f Func[T, bool]) OperatorR[T, int] {
	return func(s iter.Seq[T]) int {
		count := 0
		for v := range s {
			if f(v) {
				count++
			}
		}
		return count
	}
}
