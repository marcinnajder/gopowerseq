package seq_test

import (

	//"gotesting/powerseq/seq"

	"testing"
	//"github.com/stretchr/testify/assert"
)

// type Point struct {
// 	X, Y int
// }

func Test123(t *testing.T) {

	// items := slices.Values([]string{"a", "aa", "b", "ccc", "bb"})
	// m := seq.GroupBy(func(s string) int {
	// 	return len(s)
	// })(items)
	// assert.Equal(t, map[int][]string{1: {"a", "b"}, 2: {"aa", "bb"}, 3: {"ccc"}}, m)
}

// func TestFilter() {

// 	items := []int{1, 2, 3, 4, 5}
// 	// items := []int{1, 21, 3, 4, 51}

// 	bla0 := Filter2(func(i int) bool { return i%2 == 0 })(slices.Values(items))

// 	//bla11 := llist.Pipe4(items, slices.Values, Filter2(isEven), Map(toString), slices.Collect)

// 	bla11 := llist.Pipe4(items,
// 		slices.Values,
// 		Filter2(func(i int) bool {
// 			return i%2 == 0
// 		}),
// 		Map(func(i int) string {
// 			return strconv.Itoa(i)
// 		}),
// 		slices.Collect)

// 	if bla2, found := Find(isEven)(slices.Values(items)); found {
// 		fmt.Println(bla2)
// 	}

// 	fmt.Println(bla11)
// 	fmt.Println(bla0)

// }
