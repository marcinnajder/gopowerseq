package seq

import (
	"iter"
)

func Pipe[T, R any](val T, f func(T) R) R {
	return f(val)
}

func Pipe2[T1, T2, T3 any](val T1, f1 func(T1) T2, f2 func(T2) T3) T3 {
	return f2(f1(val))
}

func Pipe3[T1, T2, T3, T4 any](val T1, f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4) T4 {
	return f3(f2(f1(val)))
}

func Pipe4[T1, T2, T3, T4, T5 any](val T1, f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5) T5 {
	return f4(f3(f2(f1(val))))
}

type OperatorTR[T, R any] func(T) R

type OperatorR[T, R any] OperatorTR[iter.Seq[T], R]

type Operator[T, R any] OperatorTR[iter.Seq[T], iter.Seq[R]]

type Func[T, R any] func(T) R
type Func2[T1, T2, R any] func(T1, T2) R
type Func3[T1, T2, T3, R any] func(T1, T2, T3) R

func identity[T any](t T) T {
	return t
}
