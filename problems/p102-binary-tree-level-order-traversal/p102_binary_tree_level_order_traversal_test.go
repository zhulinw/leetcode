package binary_tree_level_order_traversal

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 [][]int
	}

	input := []Input{
		{
			Para1: []int{3, 9, 20, 0, 0, 15, 7},
		},
		{
			Para1: []int{1},
		},
	}

	output := []Output{
		{Para1: [][]int{{3}, {9, 20}, {15, 7}}},
		{Para1: [][]int{{1}}},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, levelOrder(tutils.IntSliceToBinaryTree(input[i].Para1)))
	}
}
