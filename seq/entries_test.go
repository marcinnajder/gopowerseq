package seq_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/seq"
	"github.com/stretchr/testify/assert"
)

func TestEntries(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}

	// order of returned entries changes randomly
	entriesMap := make(map[string]seq.Tuple[string, int])
	for t := range seq.Entries(m) {
		entriesMap[t.Item1] = t
	}
	assert.Equal(t, seq.Tuple[string, int]{"one", 1}, entriesMap["one"])
	assert.Equal(t, seq.Tuple[string, int]{"two", 2}, entriesMap["two"])
}
