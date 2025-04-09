package seq

import (
	"cmp"
	"iter"
)

func Sum[T cmp.Ordered]() OperatorR[T, T] {
	return SumFunc[T](identity)
}

func SumFunc[T any, A cmp.Ordered](f Func[T, A]) OperatorR[T, A] {
	return func(s iter.Seq[T]) A {
		var total A
		for item := range s {
			total += f(item)
		}
		return total
	}
}
