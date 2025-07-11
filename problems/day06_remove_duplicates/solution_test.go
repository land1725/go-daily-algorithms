package day06_remove_duplicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedLen   int
		expectedArray []int
	}{
		{
			name:          "empty array",
			input:         []int{},
			expectedLen:   0,
			expectedArray: []int{},
		},
		{
			name:          "single element",
			input:         []int{1},
			expectedLen:   1,
			expectedArray: []int{1},
		},
		{
			name:          "no duplicates",
			input:         []int{1, 2, 3},
			expectedLen:   3,
			expectedArray: []int{1, 2, 3},
		},
		{
			name:          "all duplicates",
			input:         []int{2, 2, 2, 2},
			expectedLen:   1,
			expectedArray: []int{2},
		},
		{
			name:          "some duplicates",
			input:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedLen:   5,
			expectedArray: []int{0, 1, 2, 3, 4},
		},
		{
			name:          "large numbers",
			input:         []int{100, 101, 101, 102, 103, 103},
			expectedLen:   4,
			expectedArray: []int{100, 101, 102, 103},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入数组的副本，因为函数会修改原数组
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			gotLen := removeDuplicates(inputCopy)

			// 检查返回的长度是否正确
			if gotLen != tt.expectedLen {
				t.Errorf("removeDuplicates() returned length = %v, want %v", gotLen, tt.expectedLen)
			}

			// 检查处理后的数组前expectedLen个元素是否正确
			if !reflect.DeepEqual(inputCopy[:gotLen], tt.expectedArray) {
				t.Errorf("after removeDuplicates(), array = %v, want %v", inputCopy[:gotLen], tt.expectedArray)
			}
		})
	}
}
