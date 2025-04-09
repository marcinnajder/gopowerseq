package seq

import (
	"iter"
)

// const set = new Set<T>();
// const set2 = new Set<T>();

// for (const s of source2) {
// 	set2.add(keySelector(s));
// }

// for (const item of source) {
// 	const key = keySelector(item);
// 	if (!set.has(key)) {
// 		set.add(key);
// 		if (!set2.has(key)) {
// 			yield item;
// 		}
// 	}
// }

func ExceptFunc[T any, K comparable](s1 iter.Seq[T], s2 iter.Seq[T], f Func[T, K]) iter.Seq[T] {
	return func(yield func(T) bool) {
		key1Set := make(map[K]struct{})
		key2Set := make(map[K]struct{})

		for item2 := range s2 {
			key2Set[f(item2)] = struct{}{}
		}

		for item1 := range s1 {
			key1 := f(item1)
			if _, ok1 := key1Set[key1]; !ok1 {
				key1Set[key1] = struct{}{}
				if _, ok2 := key2Set[key1]; !ok2 {
					if !yield(item1) {
						return
					}
				}
			}
		}
	}
}

func Except[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return ExceptFunc(s1, s2, identity[T])
}

func ExceptFuncP[T any, K comparable](s2 iter.Seq[T], f Func[T, K]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return ExceptFunc(s1, s2, f)
	}
}

func ExceptP[T comparable](s2 iter.Seq[T]) Operator[T, T] {
	return func(s1 iter.Seq[T]) iter.Seq[T] {
		return Except(s1, s2)
	}
}
