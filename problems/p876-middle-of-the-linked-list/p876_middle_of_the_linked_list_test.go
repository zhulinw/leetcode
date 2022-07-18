package middle_of_the_linked_list

import (
	"github.com/zhulinw/leetcode/testify/assert"
	"github.com/zhulinw/leetcode/testify/tutils"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 []int
	}

	input := []Input{
		{
			Para1: []int{1, 2, 3, 4, 5},
		},
		{
			Para1: []int{1, 2, 3, 4, 5, 6},
		},
	}

	output := []Output{
		{Para1: []int{3, 4, 5}},
		{Para1: []int{4, 5, 6}},
	}

	for i := range input {
		assert.Equal(t, tutils.SliceToListNode(output[i].Para1),
			middleNode2(tutils.SliceToListNode(input[i].Para1)))
	}
}
