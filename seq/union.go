package seq

import (
	"iter"

	"github.com/marcinnajder/gopowerseq/sequ"
)

func UnionFunc[T any, K comparable](s1 iter.Seq[T], s2 iter.Seq[T], f Func[T, K]) iter.Seq[T] {
	return func(yield func(T) bool) {
		keySet := make(map[K]struct{})
		for _, s := range []iter.Seq[T]{s1, s2} {
			for item := range s {
				key := f(item)
				if _, ok := keySet[key]; !ok {
					keySet[key] = struct{}{}
					if !yield(item) {
						return
					}
				}
			}
		}
	}
}

func Union[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return UnionFunc(s1, s2, sequ.Identity[T])
}

func UnionFuncP[T any, K comparable](s2 iter.Seq[T], f Func[T, K]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return UnionFunc(s1, s2, f)
	}
}

func UnionP[T comparable](s2 iter.Seq[T]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return Union(s1, s2)
	}
}
