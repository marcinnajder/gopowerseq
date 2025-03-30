package seqs

import (
	"iter"
)

func Zip[T1, T2, R any](s2 iter.Seq[T2], f Func2[T1, T2, R]) Operator[T1, R] {
	return func(s1 iter.Seq[T1]) iter.Seq[R] {
		return func(yield func(R) bool) {
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

				if !yield(f(val1, val2)) {
					return
				}
			}
		}
	}
}

// func Zip3[T1, T2, T3, R any](s2 iter.Seq[T2], s3 iter.Seq[T3], f Func3[T1, T2, T3, R]) Operator[T1, R] {
// 	return func(s1 iter.Seq[T1]) iter.Seq[R] {
// 		return func(yield func(R) bool) {
// 			next1, stop1 := iter.Pull(s1)
// 			next2, stop2 := iter.Pull(s2)
// 			next3, stop3 := iter.Pull(s3)

// 			defer stop1()
// 			defer stop2()
// 			defer stop3()

// 			for {
// 				val1, ok1 := next1()
// 				if !ok1 {
// 					return
// 				}

// 				val2, ok2 := next2()
// 				if !ok2 {
// 					return
// 				}

// 				val3, ok3 := next3()
// 				if !ok3 {
// 					return
// 				}

// 				if !yield(f(val1, val2, val3)) {
// 					return
// 				}
// 			}
// 		}
// 	}
// }
