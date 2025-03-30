package seqs_test

import (
	"testing"

	"github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seqs"
	"github.com/stretchr/testify/assert"
)

func TestCountFunc(t *testing.T) {

	assert.Equal(t, 2, seqs.CountFunc([]int{1, 2, 3, 4, 5}, shared4tests.IsEven))
	//assert.Equal(t, 0, seqs.CountFunc(isEven)(seq.Of(1, 3, 5)))
}
