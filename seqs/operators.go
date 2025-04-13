package seqs

import (
	"cmp"
	"iter"
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

// func Filter[S ~[]T, T any](s S, f seq.Func[T, bool]) iter.Seq[T] {
// 	return seq.Filter(f)(slices.Values(s))
// }

func Filter[T any](s []T, f seq.Func[T, bool]) iter.Seq[T] {
	return seq.Filter(f)(slices.Values(s))
}

func CountFunc[T any](s []T, f seq.Func[T, bool]) int {
	return seq.CountFunc(f)(slices.Values(s))
}

func CountBy[T any, K comparable](s []T, f seq.Func[T, K]) map[K]int {
	return seq.CountBy(f)(slices.Values(s))
}

func Find[T any](s []T, f seq.Func[T, bool]) (val T, index int) {
	return seq.Find(f)(slices.Values(s))
}

func All[T any](s []T, f seq.Func[T, bool]) bool {
	return seq.All(f)(slices.Values(s))
}

func Any[T any](s []T, f seq.Func[T, bool]) bool {
	return seq.Any(f)(slices.Values(s))
}

func Concat[T any](ss ...[]T) iter.Seq[T] {
	sss := Map(ss, slices.Values)
	return seq.Concat(sss...)
}

func Reduce[T any](s []T, f seq.Func2[T, T, T]) T {
	return seq.Reduce(f)(slices.Values(s))
}

func ReduceA[T, A any](s []T, f seq.Func2[A, T, A], initValue A) A {
	return seq.ReduceA(f, initValue)(slices.Values(s))
}

func Scan[T, A any](s []T, f seq.Func2[A, T, A], initValue A) iter.Seq[A] {
	return seq.Scan(f, initValue)(slices.Values(s))
}

func GroupBy[T any, K comparable](s []T, f seq.Func[T, K]) map[K][]T {
	return seq.GroupBy(f)(slices.Values(s))
}

func GroupByV[T any, K comparable, V any](s []T, keyFunc seq.Func[T, K], valueFunc seq.Func[T, V]) map[K][]V {
	return seq.GroupByV(keyFunc, valueFunc)(slices.Values(s))
}

func ToMap[T any, K comparable](s []T, keyFunc seq.Func[T, K]) map[K]T {
	return seq.ToMap(keyFunc)(slices.Values(s))
}

func ToMapV[T any, K comparable, V any](s []T, keyFunc seq.Func[T, K], valueFunc seq.Func[T, V]) map[K]V {
	return seq.ToMapV(keyFunc, valueFunc)(slices.Values(s))
}

func Skip[T any](s []T, count int) iter.Seq[T] {
	return seq.Skip[T](count)(slices.Values(s))
}

func TakeWhile[T any](s []T, f seq.Func[T, bool]) iter.Seq[T] {
	return seq.TakeWhile(f)(slices.Values(s))
}

func SkipWhile[T any](s []T, f seq.Func[T, bool]) iter.Seq[T] {
	return seq.SkipWhile(f)(slices.Values(s))
}

func DistinctFunc[T any, K comparable](s []T, f seq.Func[T, K]) iter.Seq[K] {
	return seq.DistinctFunc(f)(slices.Values(s))
}

func Distinct[T comparable](s []T) iter.Seq[T] {
	return seq.Distinct[T]()(slices.Values(s))
}

func DistinctBy[T any, K comparable](s []T, f seq.Func[T, K]) iter.Seq[T] {
	return seq.DistinctBy(f)(slices.Values(s))
}

func DistinctUntilChangedFunc[T any, K comparable](s []T, f seq.Func[T, K]) iter.Seq[T] {
	return seq.DistinctUntilChangedFunc(f)(slices.Values(s))
}

func DistinctUntilChanged[T comparable](s []T) iter.Seq[T] {
	return seq.DistinctUntilChanged[T]()(slices.Values(s))
}

func Zip[T1, T2 any](s1 []T1, s2 []T2) iter.Seq2[T1, T2] {
	return seq.Zip(slices.Values(s1), slices.Values(s2))
}

func Join[T1 any, T2 any, K comparable](s1 []T1, s2 []T2, key1 seq.Func[T1, K], key2 seq.Func[T2, K]) iter.Seq2[T1, T2] {
	return seq.Join(slices.Values(s1), slices.Values(s2), key1, key2)
}

func JoinFunc[T1 any, T2 any, K comparable, R any](s1 []T1, s2 []T2, key1 seq.Func[T1, K], key2 seq.Func[T2, K], result seq.Func2[T1, T2, R]) iter.Seq[R] {
	return seq.JoinFunc(slices.Values(s1), slices.Values(s2), key1, key2, result)
}

func UnionFunc[T any, K comparable](s1 []T, s2 []T, f seq.Func[T, K]) iter.Seq[T] {
	return seq.UnionFunc(slices.Values(s1), slices.Values(s2), f)
}

func Union[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return Union(s1, s2)
}

func IntersectFunc[T any, K comparable](s1 []T, s2 []T, f seq.Func[T, K]) iter.Seq[T] {
	return seq.IntersectFunc(slices.Values(s1), slices.Values(s2), f)
}

func Intersect[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return Intersect(s1, s2)
}

func ExceptFunc[T any, K comparable](s1 []T, s2 []T, f seq.Func[T, K]) iter.Seq[T] {
	return seq.ExceptFunc(slices.Values(s1), slices.Values(s2), f)
}

func Except[T comparable](s1 iter.Seq[T], s2 iter.Seq[T]) iter.Seq[T] {
	return Except(s1, s2)
}

func Pairwise[T any](s []T) iter.Seq2[T, T] {
	return seq.Pairwise[T]()(slices.Values(s))
}

func PairwiseFunc[T, R any](s []T, f seq.Func2[T, T, R]) iter.Seq[R] {
	return seq.PairwiseFunc(f)(slices.Values(s))
}

func Windowed[T any](s []T, size int) iter.Seq[[]T] {
	return seq.Windowed[T](size)(slices.Values(s))
}

func Chunk[T any](s []T, size int) iter.Seq[[]T] {
	return seq.Chunk[T](size)(slices.Values(s))
}

func Sum[T cmp.Ordered](s []T) T {
	return seq.Sum[T]()(slices.Values(s))
}

func SumFunc[T any, A cmp.Ordered](s []T, f seq.Func[T, A]) A {
	return seq.SumFunc(f)(slices.Values(s))
}

func Average[T seq.Number](s []T) T {
	return seq.Average[T]()(slices.Values(s))
}

func AverageFunc[T any, A seq.Number](s []T, f seq.Func[T, A]) A {
	return seq.AverageFunc(f)(slices.Values(s))
}

func Repeat[T any](s []T, count int) iter.Seq[T] {
	return seq.Repeat[T](count)(slices.Values(s))
}

func Expand[T any](s []T, f func(T) iter.Seq[T]) iter.Seq[T] {
	return seq.Expand(f)(slices.Values(s))
}

func PartitionBy[T any, V comparable](s []T, f seq.Func[T, V]) iter.Seq[[]T] {
	return seq.PartitionBy(f)(slices.Values(s))
}

func Combinations[T any](s []T, size int) iter.Seq[[]T] {
	if size > len(s) {
		return seq.Empty[[]T]()
	}
	if size == len(s) {
		return seq.Of(slices.Clone(s))
	}
	return seq.Combinations[T](size)(slices.Values(s))
}
