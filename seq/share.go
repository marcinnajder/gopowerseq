package seq

import (
	"iter"
)

func Share[T any]() Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		next, stop := iter.Pull(s)
		return pullToPush(next, stop)
	}
}

// func Share2[T any]() Operator[T, T] {
// 	return func(s iter.Seq[T]) iter.Seq[T] {

// 		subscribed := false
// 		funcs := make([]func(T) bool, 0)

// 		return func(yield func(T) bool) {
// 			if !subscribed {
// 				s(func(value T) bool {
// 					return true
// 				})
// 				subscribed = true
// 			}

// 			funcs = append(funcs, yield)

// 		}
// 	}
// }
