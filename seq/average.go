package seq

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// type Number interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 |
// 		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
// 		~float32 | ~float64
// }

func Average[T Number]() OperatorR[T, T] {
	return AverageFunc[T](identity)
}

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func AverageFunc[T any, A Number](f Func[T, A]) OperatorR[T, A] {
	return func(s iter.Seq[T]) A {
		var count, total A
		for item := range s {
			total += f(item)
			count++
		}
		if count == 0 {
			panic("seq.Average: sequence cannot be empty")
		}
		return total / count
	}
}
