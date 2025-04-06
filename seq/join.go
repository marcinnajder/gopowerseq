package seq

import (
	"iter"
)

func Bla() {

}

func Join[T1 any, T2 any, K comparable](s1 iter.Seq[T1], s2 iter.Seq[T2], key1 Func[T1, K], key2 Func[T2, K]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		m := make(map[K][]T2)

		for item2 := range s2 {
			key := key2(item2)
			m[key] = append(m[key], item2)
		}

		for item1 := range s1 {
			key := key1(item1)
			if items2, ok := m[key]; ok {
				for _, item2 := range items2 {
					if !yield(item1, item2) {
						return
					}
				}
			}
		}
	}
}

func JoinFunc[T1 any, T2 any, K comparable, R any](s1 iter.Seq[T1], s2 iter.Seq[T2], key1 Func[T1, K], key2 Func[T2, K], result Func2[T1, T2, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item1, item2 := range Join(s1, s2, key1, key2) {
			if !yield(result(item1, item2)) {
				return
			}
		}
	}
}

func JoinP[T1, T2 any, K comparable](s2 iter.Seq[T2], key1 Func[T1, K], key2 Func[T2, K]) OperatorR[T1, iter.Seq2[T1, T2]] {
	return func(s1 iter.Seq[T1]) iter.Seq2[T1, T2] {
		return Join(s1, s2, key1, key2)
	}
}

func JoinFuncP[T1 any, T2 any, K comparable, R any](s2 iter.Seq[T2], key1 Func[T1, K], key2 Func[T2, K], result Func2[T1, T2, R]) Operator[T1, R] {
	return func(s1 iter.Seq[T1]) iter.Seq[R] {
		return JoinFunc(s1, s2, key1, key2, result)
	}
}
