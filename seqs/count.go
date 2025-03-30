package seqs

import (
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func CountFunc[T any](s []T, f seq.Func[T, bool]) int {
	return seq.CountFunc(f)(slices.Values(s))
}
