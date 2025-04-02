package seqs

import (
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func Any[T any](s []T, f seq.Func[T, bool]) bool {
	return seq.Any(f)(slices.Values(s))
}
