package day02_single_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "常规用例 - 唯一数字在中间",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "常规用例 - 唯一数字在开头",
			nums: []int{1, 2, 2, 3, 3},
			want: 1,
		},
		{
			name: "常规用例 - 唯一数字在结尾",
			nums: []int{2, 2, 3, 3, 1},
			want: 1,
		},
		{
			name: "边界测试 - 只有一个元素",
			nums: []int{42},
			want: 42,
		},
		{
			name: "边界测试 - 大数字",
			nums: []int{1000000, 2, 1000000},
			want: 2,
		},
		{
			name: "异常测试 - 无唯一数字（默认返回-1）",
			nums: []int{1, 1, 2, 2},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			assert.Equal(t, tt.want, got, "输入: %v", tt.nums)
		})
	}
}
