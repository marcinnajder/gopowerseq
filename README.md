## Documentation

**gopowerseq** library provides a set of standard operators working with sequences in Go. [go1.23](https://go.dev/doc/go1.23#iterators) (2024-08-13) introduced [iterators](https://go.dev/blog/range-functions) (`iter.Seq[T]` interface) that can be treated as a lazy sequences known from other programming languages like C#, JS, F#, Clojure, Python, ... . This library is a port of the JS counterpart [powerseq](https://github.com/marcinnajder/powerseq).

```go
isEven := func(x int) bool { return x%2 == 0 }

var numbersSlice []int = []int{1, 2, 3, 4, 5}
var numbersIter iter.Seq[int] = seq.Of(1, 2, 3, 4, 5) // slices.Values(numbersSlice)

// 'seq' package - calling single operator taking input of type 'iter.Seq[T]'
var evenNumbers iter.Seq[int] = seq.Filter(isEven)(numbersIter)
for n := range evenNumbers {
	fmt.Println(n) // 2, 4
}

// 'seq' package - chaining many operators
var items []int = seq.Pipe3(
	seq.Range(0, math.MaxInt64),
	seq.Filter(func(x int) bool { return x%2 == 0 }),
	seq.Take[int](5),
	seq.ToSlice[int]())
fmt.Println(items) // [0 2 4 6 8]

// 'seqs' package - calling single operator taking input of type 'slice'
var chars = []string{"a", "b", "c"}
for n, c := range seqs.Zip(numbersSlice, chars) {
	fmt.Println(n, "-", c) // "1 - a", "2 - b", "3 - c"
}
```

- operators from `seq` package take `iter.Seq[]` and return functions `func (iter.Seq[]) iter.Seq[]`, there are "curried" so we can `Pipe` many operators easily
- operators from `seqs` package take slices `[]T` and return sequences `iter.Seq[]`
- for some operators a special counterparts ending with `P` are provided (`ConcatP`,`ExceptP`,`InterleaveP`,`IntersectP`,`JoinP`,`SequenceEqualP`, `UnionP`, `ZipP`), those functions must be used inside `Pipe(..., opp())`, so we call `Concat([1,2,3], [4,5,6])` or `Pipe([1,2,3], ConcatP([4,5,6]) )`
- read [Programming with sequences](https://marcinnajder.github.io/2022/11/02/programming-with-sequences-part-1-introduction-to-powerseq.html) series to better understand iterators and generators in JS, it also presents how to implement powerseq operators from scratch

### operators

- creation - Of, Empty, Entries, Range, Repeat, RepeatValue
- filtering - Filter, Skip, SkipWhile, Take, TakeWhile
- mapping - Map, FlatMap, FlatMapR
- partitioning - Pairwise, Windowed, PartitionBy, Chunk, Combinations
- merging - Join, Zip, Interleave, Interpose
- grouping - GroupBy, GroupByV, CountBy
- set - Except, Intersect, Union, Distinct, DistinctBy, DistinctUntilChanged
- conversion - ToSlice, ToMap, ToMapV, ToTuples
- aggregation - Reduce, ReduceA, Average, Count, Sum
- quantifiers - All, Any
- concatenation - Concat
- equality - SequenceEqual
- element - First
- misc - Memoize, Share, Expand, Scan

```go
// helpers
func IsEven(n int) bool       { return n%2 == 0 }
func Plus(a, b int) int       { return a + b }
func GetLen(next string) int  { return len(next) }
func Identity[T any](val T) T { return val }

// operators
seq.All(IsEven)(seq.Of(2, 4, 6))  // true
seq.All(IsEven)(seq.Of(2, 4, 5, 6)) // false

seq.Any(IsEven)(seq.Of(1, 3, 5))  // false
seq.Any(IsEven)(seq.Of(1, 3, 4, 5)) // true

seq.Average[int]()(seq.Of(5, 10, 15)) // 10
seq.AverageFunc(GetLen)(seq.Of("a", "bb", "ccc")) // 2

seq.Chunk[int](2)(seq.Of(5, 10, 15, 50, 100)) // [5 10], [15 50], [100]

seq.Combinations[int](2)(seq.Of(5, 10, 15)) // [5 10], [5 15], [10 15]

seq.Concat(seq.Of(1, 2), seq.Of[int](), seq.Of(3)) // 1, 2, 3

seq.Count[int]()(seq.Of(1, 2, 3, 4)) // 4
seq.CountFunc(IsEven)(seq.Of(1, 2, 3, 4)) // 2

seq.CountBy(GetLen)(seq.Of("a", "bb", "c", "dddd")) // map[1:2 2:1 4:1]

seq.Distinct[int]()(seq.Of(1, 2, 1, 3, 4, 2)) // 1, 2, 3, 4
seq.DistinctFunc(GetLen)(seq.Of("a", "bb", "c", "dddd")) // 1, 2, 4
seq.DistinctBy(GetLen)(seq.Of("a", "bb", "c", "dddd")) // a, bb, dddd
seq.DistinctUntilChanged[int]()(seq.Of(1, 2, 2, 3, 3, 1, 4)) // 1, 2, 3, 1, 4
seq.DistinctUntilChangedFunc(GetLen)(seq.Of("a", "b", "cc", "d", "e")) // a, cc, d

seq.Empty[int]()

seq.Entries(map[int]string{1: "one", 2: "two"}) // {2 two}, {1 one}

seq.Except(seq.Of(1, 2, 2, 3, 3), seq.Of(3, 3, 4, 4)) // 1, 2

expanded := seq.Expand(func(str string) iter.Seq[string] {
	if len(str) > 1 {
		return slices.Values(strings.Split(str, ""))
	} else {
		return seq.Empty[string]()
	}
})(seq.Of("a", "xy", "b")) // "a", "xy", "b", "x", "y"

seq.Filter(IsEven)(seq.Of(1, 2, 2, 4, 5, 7, 8, 10)) // 2, 2, 4, 8, 10

seq.Find(IsEven)(seq.Of(1, 3, 5, 6, 5, 4)) // val: 6, index: 3
seq.Find(IsEven)(seq.Of(1, 3, 5, 5)) // val: 0, index: -1

flatten := seq.FlatMap(func(n int) iter.Seq[string] {
	return slices.Values(slices.Repeat([]string{strconv.Itoa(n)}, n))
})(seq.Of(1, 2, 3)) // "1", "2", "2", "3", "3", "3"

seq.GroupBy(GetLen)(seq.Of("a", "bb", "c")) // map[1:[a c] 2:[bb]]
seq.GroupByV(GetLen, strings.ToUpper)(seq.Of("a", "bb", "c")) // map[1:[A C] 2:[BB]]

seq.Interleave(seq.Of(1, 2, 3), seq.Of(5, 10), seq.Of(-1, -2)) // 1, 5, -1, 2, 10, -2

seq.Interpose(0)(seq.Of(1, 2, 3)) // 1, 0, 2, 0, 3

seq.Intersect(seq.Of(1, 2, 3), seq.Of(2, 3, 4))  //  2, 3
seq.IntersectFunc(seq.Of("a", "bb"), seq.Of("c", "ddd", "ee"), GetLen) // "c", "ee"

seq.Join(seq.Of(1, 2, 3), seq.Of("a", "bbb", "ccc"), Identity, GetLen) // {1 a}, {3 bbb}, {3 ccc}

seq.JoinFunc(seq.Of(1, 2, 3), seq.Of("a", "bbb", "ccc"), Identity, GetLen,
	func(n int, s string) string {
		return fmt.Sprint(n, "-", s)
	}) // 1-a, 3-bbb, 3-ccc

seq.Map(IsEven)(seq.Of(1, 2, 3)) // false, true, false

myIterState := 0
myIter := func(yield func(int) bool) {
	for {
		if !yield(myIterState) {
			return
		}
		myIterState++
	}
}
memoized := seq.Take[int](3)(seq.Memoize[int]()(myIter))
memoized // 0, 1, 2
memoized // 0, 1, 2 (without calling Memoize: 2, 3, 4)

seq.Of(1, 2, 3) // 1, 2, 3

seq.Pairwise[int]()(seq.Of(1, 2, 3, 4))  // {1 2}, {2 3}, {3 4}
seq.PairwiseFunc(func(a, b int) []int { return []int{a, b} })(seq.Of(1, 2, 3, 4)) // [1 2], [2 3], [3 4]

seq.PartitionBy(IsEven)(seq.Of(1, 3, 4, 6, 7, 8)) // [1 3], [4 6], [7], [8]

seq.Range(8, 4) // 8, 9, 10, 11

seq.Reduce(Plus)(seq.Of(1, 2, 3)) // 6
seq.ReduceA(func(s string, b int) string { return fmt.Sprint(s, "-", b) }, "")(seq.Of(1, 2, 3)) // "-1-2-3"

seq.Repeat[int](3)(seq.Of(1, 2))  // 1, 2, 1, 2, 1, 2
seq.Take[int](5)(seq.Repeat[int](-1)(seq.Of(1, 2)))  // 1, 2, 1, 2, 1

seq.RepeatValue(12, 3)  // 12, 12, 12
seq.Take[int](4)(seq.RepeatValue(12, -1)))  // 12, 12, 12, 12

seq.Scan(func(s string, b int) string { return fmt.Sprint(s, "-", b) }, "")(seq.Of(1, 2, 3))
// "", "-1", "-1-2", "-1-2-3"

seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2, 3))  // true
seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2))  // false

shared := seq.Share[int]()(seq.Of(1, 2, 3))

seq.Take[int](2)(shared) // 1, 2
seq.Take[int](2)(shared) // 3 (without calling Share: 1, 2)

seq.Skip[int](2)(seq.Of(1, 2, 3, 4)) // 3, 4

seq.SkipWhile(IsEven)(seq.Of(2, 4, 3, 4, 7, 8)) // 3, 4, 7, 8

seq.Sum[int]()(seq.Of(1, 2, 3))  // 6
seq.SumFunc(GetLen)(seq.Of("a", "bb", "c"))     // 4

seq.Take[int](2)(seq.Of(1, 2, 3, 4)) // 1, 2

seq.TakeWhile(IsEven)(seq.Of(2, 4, 3, 4, 7, 8)) // 2, 4

seq.ToMap(GetLen)(seq.Of("a", "bb", "c", "dddd"))  // map[1:c 2:bb 4:dddd]
seq.ToMapV(GetLen, strings.ToUpper)(seq.Of("a", "bb", "c", "dddd")) // map[1:C 2:BB 4:DDDD]

seq.ToSlice[int]()(seq.Of(1, 2, 3))  // [1 2 3]

seq.ToTuples[int, string]()(maps.All(map[int]string{1: "one", 2: "two"}))  // {1 one}, {2 two}

seq.Union(seq.Of(1, 2, 3), seq.Of(2, 3, 4, 4))  //  1, 2, 3, 4
seq.UnionFunc(seq.Of("a", "bb"), seq.Of("c", "ddd", "ee"), GetLen)  // "a", "bb", "ddd"

seq.Windowed[int](3)(seq.Of(1, 222, 3, 44, 5)) // [1 222 3], [222 3 44], [3 44 5]

seq.Zip(seq.Of(1, 2, 3), seq.Of("a", "b"))  // {1 a}, {2 b}
seq.ZipFunc(seq.Of(1, 2, 3), seq.Of(10, 100), Plus) // 11, 102
```
