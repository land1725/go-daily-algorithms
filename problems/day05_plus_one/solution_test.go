package day05_plus_one

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "no carry",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "single digit no carry",
			input:    []int{5},
			expected: []int{6},
		},
		{
			name:     "carry at last digit",
			input:    []int{1, 9, 9},
			expected: []int{2, 0, 0},
		},
		{
			name:     "multiple carries",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "empty slice", // 边界情况
			input:    []int{},
			expected: []int{1},
		},
		{
			name:     "all zeros", // 非常规输入测试
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
