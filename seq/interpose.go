package seq

import (
	"iter"
)

func Interpose[T any](sep T) Operator[T, T] {
	return func(s iter.Seq[T]) iter.Seq[T] {
		return Pipe2(s,
			FlatMap(func(item T) iter.Seq[T] { return Of(sep, item) }),
			Skip[T](1))
	}
}
