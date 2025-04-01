package seqs

import (
	"slices"

	"github.com/marcinnajder/gopowerseq/seq"
)

func Find[T any](s []T, f seq.Func[T, bool]) (val T, index int) {
	return seq.Find(f)(slices.Values(s))
}
