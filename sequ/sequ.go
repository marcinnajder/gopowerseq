package sequ

func If[T any](cond bool, truevalue T, falsevalue T) T {
	if cond {
		return truevalue
	} else {
		return falsevalue
	}
}

func Identity[T any](t T) T {
	return t
}

func Zero[T any]() T {
	var zero T
	return zero
}
