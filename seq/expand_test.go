package seq_test

import (
	"iter"
	"slices"
	"testing"

	. "github.com/marcinnajder/gopowerseq/internal/shared4tests"
	"github.com/marcinnajder/gopowerseq/seq"
)

type TreeNode struct {
	name     string
	children []TreeNode
}

func TestExpand(t *testing.T) {
	numbers1 := seq.Expand(func(x int) iter.Seq[int] { return seq.If(x > 8, seq.Empty[int](), seq.Of(x*2)) })(seq.Of(1))
	Assert_EqualSeq(t, seq.Of(1, 2, 4, 8, 16), numbers1)

	numbers2 := seq.Expand(func(x int) iter.Seq[int] { return seq.If(x > 8, nil, seq.Of(x*2)) })(seq.Of(1))
	Assert_EqualSeq(t, seq.Of(1, 2, 4, 8, 16), numbers2)

	empty := []TreeNode{}
	tree := []TreeNode{{"a", []TreeNode{{"d", []TreeNode{{"dd", empty}}}, {"w", []TreeNode{}}}}, {"x", empty}}

	nodes := seq.Expand(func(node TreeNode) iter.Seq[TreeNode] {
		return slices.Values(node.children)
	})(slices.Values(tree))
	names := seq.Map(func(n TreeNode) string { return n.name })(nodes)
	Assert_EqualSeq(t, seq.Of("a", "x", "d", "w", "dd"), names)
}
