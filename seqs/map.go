package seqs

import (
	"github.com/marcinnajder/gopowerseq/seq"
)

func Map[T, R any](s []T, f seq.Func[T, R]) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}
