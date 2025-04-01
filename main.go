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
