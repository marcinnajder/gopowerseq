package seq

import (
	"iter"
	"slices"
)

// function _combinations<T>(source: Iterable<T>, size: number) {
//     return wrapInIterable(function* () {
//         if (size <= 0) {
//             return;
//         }

//         if (Array.isArray(source)) {
//             if (size > source.length) {
//                 return;
//             }
//             if (size === source.length) {
//                 yield [...source]; // copy
//                 return;
//             }
//         }

//         const allTuples: T[][] = [[]];
//         for (const item of source) {
//             const newTuples: T[][] = [];
//             for (const tuple of allTuples) {
//                 const newTuple = [...tuple, item];
//                 if (newTuple.length === size) {
//                     yield newTuple;
//                 } else {
//                     newTuples.push(newTuple);
//                 }
//             }
//             allTuples.push(...newTuples);
//         }
//     });
// }

func Combinations[T any](size int) OperatorR[T, iter.Seq[[]T]] {
	if size <= 0 {
		panic("seq.Combinations: 'size' must be positive")
	}

	return func(s iter.Seq[T]) iter.Seq[[]T] {
		return func(yield func([]T) bool) {
			allTuples := [][]T{{}}
			for item := range s {
				newTuples := make([][]T, 0)
				for _, tuple := range allTuples {
					newTuple := append(slices.Clone(tuple), item)
					if len(newTuple) == size {
						if !yield(newTuple) {
							return
						}
					} else {
						newTuples = append(newTuples, newTuple)
					}
				}
				allTuples = append(allTuples, newTuples...)
			}
		}
	}
}
