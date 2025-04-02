package seq

import (
	"iter"
)

func Any[T any](f Func[T, bool]) OperatorR[T, bool] {
	return func(s iter.Seq[T]) bool {
		for item := range s {
			if f(item) {
				return true
			}
		}
		return false
	}
}
