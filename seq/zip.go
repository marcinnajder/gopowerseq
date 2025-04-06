package seq

import (
	"iter"
)

func Zip[T1, T2 any](s1 iter.Seq[T1], s2 iter.Seq[T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		next1, stop1 := iter.Pull(s1)
		next2, stop2 := iter.Pull(s2)

		defer stop1()
		defer stop2()

		for {
			val1, ok1 := next1()
			if !ok1 {
				return
			}

			val2, ok2 := next2()
			if !ok2 {
				return
			}

			if !yield(val1, val2) {
				return
			}
		}
	}
}

func ZipFunc[T1, T2, R any](s1 iter.Seq[T1], s2 iter.Seq[T2], f Func2[T1, T2, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item1, item2 := range Zip(s1, s2) {
			if !yield(f(item1, item2)) {
				return
			}
		}
	}
}

func Zip3Func[T1, T2, T3, R any](s1 iter.Seq[T1], s2 iter.Seq[T2], s3 iter.Seq[T3], f Func3[T1, T2, T3, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		next1, stop1 := iter.Pull(s1)
		next2, stop2 := iter.Pull(s2)
		next3, stop3 := iter.Pull(s3)

		defer stop1()
		defer stop2()
		defer stop3()

		for {
			val1, ok1 := next1()
			if !ok1 {
				return
			}

			val2, ok2 := next2()
			if !ok2 {
				return
			}

			val3, ok3 := next3()
			if !ok3 {
				return
			}

			if !yield(f(val1, val2, val3)) {
				return
			}
		}
	}
}

func ZipP[T1, T2 any](s2 iter.Seq[T2]) OperatorR[T1, iter.Seq2[T1, T2]] {
	return func(s1 iter.Seq[T1]) iter.Seq2[T1, T2] {
		return Zip(s1, s2)
	}
}

func ZipFuncP[T1, T2, R any](s2 iter.Seq[T2], f Func2[T1, T2, R]) Operator[T1, R] {
	return func(s1 iter.Seq[T1]) iter.Seq[R] {
		return ZipFunc(s1, s2, f)
	}
}

func Zip3FuncP[T1, T2, T3, R any](s2 iter.Seq[T2], s3 iter.Seq[T3], f Func3[T1, T2, T3, R]) Operator[T1, R] {
	return func(s1 iter.Seq[T1]) iter.Seq[R] {
		return Zip3Func(s1, s2, s3, f)
	}
}
