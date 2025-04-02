package seqs

import (
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func All[T any](s []T, f seq.Func[T, bool]) bool {
	return seq.All(f)(slices.Values(s))
}
