// go test ./...

package main

import (
	"fmt"
	"math"

	"github.com/marcinnajder/gopowerseq/seq"
)

func main() {

	powerSeqIntro()
	//fmt.Println(q)

	// q := seq.Pipe2(
	// 	seq.Range(1, math.MaxInt64),
	// 	seq.Filter(func(x int) bool { return x%2 == 0 }),
	// 	seq.Take(5))
	// seq.Range(
	// fmt.Println("asasd")
	// bla := seq.Of(1, 2, 3)
	// fmt.Println("asasd", bla)
}

func powerSeqIntro() {
	for item := range seq.Filter(func(x int) bool { return x%2 == 0 })(seq.Of(1, 2, 3, 4, 5)) {
		fmt.Println(item)
	}

	items := seq.Pipe3(
		seq.Range(0, math.MaxInt64),
		seq.Filter(func(x int) bool { return x%2 == 0 }),
		seq.Take[int](5),
		seq.ToSlice)

	fmt.Println(items)
}

// 	import { pipe, range, filter, take, toarray } from "powerseq"; // npm install powerseq

// // calling one operator
// for(var item of filter([1,2,3,4,5], x => x % 2 === 0)){
//     console.log(item);
// }

// // chaining many operators
// const items = pipe(range(1, Number.MAX_VALUE), filter(x => x % 2 === 0), take(5), toarray());
// console.log(items);

// items := []int{123, 123, 123, 123, 123}
// fiteredItems := make([]int, 0)
// for _, item := range items {
// 	if item > 10 { // Filter!!
// 		fiteredItems = append(fiteredItems, item*1000) // Map!!
// 	}
// }

// aaaaa := seq.Filter(func(item int) bool {
// 	return item > 10
// })(slices.Values(items))

// // aaaaa := seqs.Map(items, func(item int) string {
// // 	return string(item)
// // })

// fiteredItems2 := seq.Pipe3(
// 	slices.Values(items),
// 	seq.Filter(func(item int) bool {
// 		return item > 10
// 	}),
// 	seq.Map(func(item int) int {
// 		return item * 100
// 	}),
// 	seq.ToSlice)

// _ = fiteredItems
// _ = fiteredItems2

// if val, i := seqs.Find(items, func(i int) bool { return i > 10 }); i != -1 {
// 	fmt.Println(val, i)
// }

// fiteredItems2 := seq.Filter(func(item int) bool {
// 	return item > 10
// })(slices.Values(items))

// for val1, val2 := range seqs.Zip__([]int{1, 123, 13}, []string{"asd", " adads"}) {

// }
