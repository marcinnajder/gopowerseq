package examples

import (
	"fmt"
	"iter"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/marcinnajder/gopowerseq/seqs"
)

func DocIntro() {

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
}

// helpers
func IsEven(n int) bool       { return n%2 == 0 }
func Plus(a, b int) int       { return a + b }
func GetLen(next string) int  { return len(next) }
func Identity[T any](val T) T { return val }

func DocOperators() {
	// operators
	printValue("All", seq.All(IsEven)(seq.Of(2, 4, 6)))    // true
	printValue("All", seq.All(IsEven)(seq.Of(2, 4, 5, 6))) // false

	printValue("Any", seq.Any(IsEven)(seq.Of(1, 3, 5)))    // false
	printValue("Any", seq.Any(IsEven)(seq.Of(1, 3, 4, 5))) // true

	printValue("Average", seq.Average[int]()(seq.Of(5, 10, 15)))                 // 10
	printValue("AverageFunc", seq.AverageFunc(GetLen)(seq.Of("a", "bb", "ccc"))) // 2

	printSeq("Chunk", seq.Chunk[int](2)(seq.Of(5, 10, 15, 50, 100))) // [5 10], [15 50], [100]

	printSeq("Combinations", seq.Combinations[int](2)(seq.Of(5, 10, 15))) // [5 10], [5 15], [10 15]

	printSeq("Concat", seq.Concat(seq.Of(1, 2), seq.Of[int](), seq.Of(3))) // 1, 2, 3

	printValue("Count", seq.Count[int]()(seq.Of(1, 2, 3, 4)))          // 4
	printValue("CountFunc", seq.CountFunc(IsEven)(seq.Of(1, 2, 3, 4))) // 2

	printValue("CountFunc", seq.CountBy(GetLen)(seq.Of("a", "bb", "c", "dddd"))) // map[1:2 2:1 4:1]

	printSeq("Distinct", seq.Distinct[int]()(seq.Of(1, 2, 1, 3, 4, 2)))                                          // 1, 2, 3, 4
	printSeq("DistinctFunc", seq.DistinctFunc(GetLen)(seq.Of("a", "bb", "c", "dddd")))                           // 1, 2, 4
	printSeq("DistinctBy", seq.DistinctBy(GetLen)(seq.Of("a", "bb", "c", "dddd")))                               // a, bb, dddd
	printSeq("DistinctUntilChanged", seq.DistinctUntilChanged[int]()(seq.Of(1, 2, 2, 3, 3, 1, 4)))               // 1, 2, 3, 1, 4
	printSeq("DistinctUntilChangedFunc", seq.DistinctUntilChangedFunc(GetLen)(seq.Of("a", "b", "cc", "d", "e"))) // a, cc, d

	printSeq("Empty", seq.Empty[int]())

	printSeq("Entries", seq.Entries(map[int]string{1: "one", 2: "two"})) // {2 two}, {1 one}

	printSeq("Except", seq.Except(seq.Of(1, 2, 2, 3, 3), seq.Of(3, 3, 4, 4))) // 1, 2

	expanded := seq.Expand(func(str string) iter.Seq[string] {
		if len(str) > 1 {
			return slices.Values(strings.Split(str, ""))
		} else {
			return seq.Empty[string]()
		}
	})(seq.Of("a", "xy", "b"))
	printSeq("Expand", expanded) // "a", "xy", "b", "x", "y"

	printSeq("Filter", seq.Filter(IsEven)(seq.Of(1, 2, 2, 4, 5, 7, 8, 10))) // 2, 2, 4, 8, 10

	foundVal1, foundIndex1 := seq.Find(IsEven)(seq.Of(1, 3, 5, 6, 5, 4))
	printValue2("Find", foundVal1, foundIndex1) // val: 6, index: 3
	foundVal2, foundIndex2 := seq.Find(IsEven)(seq.Of(1, 3, 5, 5))
	printValue2("Find", foundVal2, foundIndex2) // val: 0, index: -1

	flatten := seq.FlatMap(func(n int) iter.Seq[string] {
		return slices.Values(slices.Repeat([]string{strconv.Itoa(n)}, n))
	})(seq.Of(1, 2, 3))
	printSeq("FlatMap", flatten) // "1", "2", "2", "3", "3", "3"

	printValue("GroupBy", seq.GroupBy(GetLen)(seq.Of("a", "bb", "c")))                    // map[1:[a c] 2:[bb]]
	printValue("GroupByV", seq.GroupByV(GetLen, strings.ToUpper)(seq.Of("a", "bb", "c"))) // map[1:[A C] 2:[BB]]

	printSeq("Interleave", seq.Interleave(seq.Of(1, 2, 3), seq.Of(5, 10), seq.Of(-1, -2))) // 1, 5, -1, 2, 10, -2

	printSeq("Interpose", seq.Interpose(0)(seq.Of(1, 2, 3))) // 1, 0, 2, 0, 3

	printSeq("Intersect", seq.Intersect(seq.Of(1, 2, 3), seq.Of(2, 3, 4)))                            //  2, 3
	printSeq("IntersectFunc", seq.IntersectFunc(seq.Of("a", "bb"), seq.Of("c", "ddd", "ee"), GetLen)) // "c", "ee"

	joined1 := seq.Join(seq.Of(1, 2, 3), seq.Of("a", "bbb", "ccc"), Identity, GetLen)
	printSeq2("Join", joined1) // {1 a}, {3 bbb}, {3 ccc}
	joined2 := seq.JoinFunc(seq.Of(1, 2, 3), seq.Of("a", "bbb", "ccc"), Identity, GetLen,
		func(n int, s string) string {
			return fmt.Sprint(n, "-", s)
		})
	printSeq("Join", joined2) // 1-a, 3-bbb, 3-ccc

	printSeq("Map", seq.Map(IsEven)(seq.Of(1, 2, 3))) // false, true, false

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
	//memoized := seq.Take[int](3)(myIter)
	printSeq("Memoize", memoized) // 0, 1, 2
	printSeq("Memoize", memoized) // 0, 1, 2 (without calling Memoize: 2, 3, 4)

	printSeq("Of", seq.Of(1, 2, 3)) // 1, 2, 3

	printSeq2("Pairwise", seq.Pairwise[int]()(seq.Of(1, 2, 3, 4)))                                              // {1 2}, {2 3}, {3 4}
	printSeq("PairwiseFunc", seq.PairwiseFunc(func(a, b int) []int { return []int{a, b} })(seq.Of(1, 2, 3, 4))) // [1 2], [2 3], [3 4]

	printSeq("PartitionBy", seq.PartitionBy(IsEven)(seq.Of(1, 3, 4, 6, 7, 8))) // [1 3], [4 6], [7], [8]

	printSeq("Range", seq.Range(8, 4)) // 8, 9, 10, 11

	printValue("Range", seq.Reduce(Plus)(seq.Of(1, 2, 3)))                                                                // 6
	printValue("RangeA", seq.ReduceA(func(s string, b int) string { return fmt.Sprint(s, "-", b) }, "")(seq.Of(1, 2, 3))) // "-1-2-3"

	printSeq("Repeat", seq.Repeat[int](3)(seq.Of(1, 2)))                    // 1, 2, 1, 2, 1, 2
	printSeq("Repeat", seq.Take[int](5)(seq.Repeat[int](-1)(seq.Of(1, 2)))) // 1, 2, 1, 2, 1

	printSeq("RepeatValue", seq.RepeatValue(12, 3))                    // 12, 12, 12
	printSeq("RepeatValue", seq.Take[int](4)(seq.RepeatValue(12, -1))) // 12, 12, 12, 12

	printSeq("Scan", seq.Scan(func(s string, b int) string { return fmt.Sprint(s, "-", b) }, "")(seq.Of(1, 2, 3))) // "", "-1", "-1-2", "-1-2-3"

	printValue("SequenceEqual", seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2, 3))) // true
	printValue("SequenceEqual", seq.SequenceEqual(seq.Of(1, 2, 3), seq.Of(1, 2)))    // false

	shared := seq.Share[int]()(seq.Of(1, 2, 3))
	//shared := seq.Of(1, 2, 3)
	printSeq("Share", seq.Take[int](2)(shared)) // 1, 2
	printSeq("Share", seq.Take[int](2)(shared)) // 3 (without calling Share: 1, 2)

	printSeq("Skip", seq.Skip[int](2)(seq.Of(1, 2, 3, 4))) // 3, 4

	printSeq("SkipWhile", seq.SkipWhile(IsEven)(seq.Of(2, 4, 3, 4, 7, 8))) // 3, 4, 7, 8

	printValue("Sum", seq.Sum[int]()(seq.Of(1, 2, 3)))                 // 6
	printValue("SumFunc", seq.SumFunc(GetLen)(seq.Of("a", "bb", "c"))) // 4

	printSeq("Take", seq.Take[int](2)(seq.Of(1, 2, 3, 4))) // 1, 2

	printSeq("TakeWhile", seq.TakeWhile(IsEven)(seq.Of(2, 4, 3, 4, 7, 8))) // 2, 4

	printValue("ToMap", seq.ToMap(GetLen)(seq.Of("a", "bb", "c", "dddd")))                    // map[1:c 2:bb 4:dddd]
	printValue("ToMapV", seq.ToMapV(GetLen, strings.ToUpper)(seq.Of("a", "bb", "c", "dddd"))) // map[1:C 2:BB 4:DDDD]

	printValue("ToSlice", seq.ToSlice[int]()(seq.Of(1, 2, 3))) // [1 2 3]

	printSeq("ToTuples", seq.ToTuples[int, string]()(maps.All(map[int]string{1: "one", 2: "two"}))) // {1 one}, {2 two}

	printSeq("Union", seq.Union(seq.Of(1, 2, 3), seq.Of(2, 3, 4, 4)))                       //  1, 2, 3, 4
	printSeq("UnionBy", seq.UnionFunc(seq.Of("a", "bb"), seq.Of("c", "ddd", "ee"), GetLen)) //  "a", "bb", "ddd"

	printSeq("Windowed", seq.Windowed[int](3)(seq.Of(1, 222, 3, 44, 5))) // [1 222 3], [222 3 44], [3 44 5]

	printSeq2("Zip", seq.Zip(seq.Of(1, 2, 3), seq.Of("a", "b")))         // {1 a}, {2 b}
	printSeq("Zip", seq.ZipFunc(seq.Of(1, 2, 3), seq.Of(10, 100), Plus)) // 11, 102
}

func ToString[T any](val T) string {
	return fmt.Sprintf("%v", val)
}

func printSeq[T any](text string, s iter.Seq[T]) {
	items := seq.Pipe2(
		s,
		seq.Map(ToString[T]),
		seq.ToSlice[string]())
	fmt.Printf("%s: %s \n", text, strings.Join(items, ", "))
}

func printValue[T any](text string, val T) {
	fmt.Printf("%s: %v \n", text, val)
}

func printValue2[T1, T2 any](text string, val1 T1, val2 T2) {
	fmt.Printf("%s: %v, %v \n", text, val1, val2)
}

func printSeq2[T1, T2 any](text string, s iter.Seq2[T1, T2]) {
	items := seq.ToTuples[T1, T2]()(s)
	printSeq(text, items)
}
