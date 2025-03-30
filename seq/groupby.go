package seq

func GroupBy[T any, K comparable](f Func[T, K]) OperatorR[T, map[K][]T] {
	return GroupValueBy(f, identity)
}
