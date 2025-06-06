package seq

import (
	"iter"
)

type Func[T, R any] func(T) R
type Func2[T1, T2, R any] func(T1, T2) R
type Func3[T1, T2, T3, R any] func(T1, T2, T3) R

// type inference does work correctly when helpers type 'Func' is used instead of typ 'func(T) R'

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

func Pipe5[T1, T2, T3, T4, T5, T6 any](val T1, f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5, f5 func(T5) T6) T6 {
	return f5(f4(f3(f2(f1(val)))))
}

func Pipe6[T1, T2, T3, T4, T5, T6, T7 any](val T1, f1 func(T1) T2, f2 func(T2) T3, f3 func(T3) T4, f4 func(T4) T5, f5 func(T5) T6, f6 func(T6) T7) T7 {
	return f6(f5(f4(f3(f2(f1(val))))))
}

type OperatorTR[T, R any] func(T) R

type OperatorR[T, R any] OperatorTR[iter.Seq[T], R]

type Operator[T, R any] OperatorTR[iter.Seq[T], iter.Seq[R]]

type Tuple[T1, T2 any] struct {
	Item1 T1
	Item2 T2
}

type Tuple3[T1, T2, T3 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
}

type Tuple4[T1, T2, T3, T4 any] struct {
	Item1 T1
	Item2 T2
	Item3 T3
	Item4 T4
}

func getIterators[T any](ss []iter.Seq[T]) (nextFuncs []func() (T, bool), stopAll func()) {
	nextFuncs = make([]func() (T, bool), len(ss))
	stopFuncs := make([]func(), len(ss))

	for i, s := range ss {
		next, stop := iter.Pull(s)
		nextFuncs[i] = next
		stopFuncs[i] = stop
	}

	stopAll = func() {
		for _, stop := range stopFuncs {
			stop()
		}
	}

	return nextFuncs, stopAll
}
