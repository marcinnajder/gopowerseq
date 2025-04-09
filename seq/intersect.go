package seq

import (
	"iter"
)

func IntersectFunc[T any, K comparable](s1 iter.Seq[T], s2 iter.Seq[T], f Func[T, K]) iter.Seq[T] {
	return func(yield func(T) bool) {
		key1Set := make(map[K]struct{})
		key2Set := make(map[K]struct{})

		for item1 := range s1 {
			key1Set[f(item1)] = struct{}{}
		}

		for item2 := range s2 {
			key2 := f(item2)
			if _, ok1 := key1Set[key2]; ok1 {
				if _, ok2 := key2Set[key2]; !ok2 {
					key2Set[key2] = struct{}{}
					if !yield(item2) {
						return
					}
				}
			}
		}
	}
}

func Intersect[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return IntersectFunc(s1, s2, identity[T])
}

func IntersectFuncP[T any, K comparable](s2 iter.Seq[T], f Func[T, K]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return IntersectFunc(s1, s2, f)
	}
}

func IntersectP[T comparable](s2 iter.Seq[T]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return Union(s1, s2)
	}
}
