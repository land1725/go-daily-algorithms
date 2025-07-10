package day01_two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"case1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"case2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"case3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}
