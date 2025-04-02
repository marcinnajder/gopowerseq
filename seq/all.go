package seq

import (
	"iter"
)

func All[T any](f Func[T, bool]) OperatorR[T, bool] {
	return func(s iter.Seq[T]) bool {
		for item := range s {
			if !f(item) {
				return false
			}
		}
		return true
	}
}
