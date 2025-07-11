package day07_is_palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		// 正数回文测试
		{
			name:     "single digit positive",
			input:    5,
			expected: true,
		},
		{
			name:     "even digits palindrome",
			input:    1221,
			expected: true,
		},
		{
			name:     "odd digits palindrome",
			input:    12321,
			expected: true,
		},
		{
			name:     "even digits non-palindrome",
			input:    1234,
			expected: false,
		},
		{
			name:     "odd digits non-palindrome",
			input:    12345,
			expected: false,
		},

		// 负数测试
		{
			name:     "negative number",
			input:    -121,
			expected: false,
		},
		{
			name:     "negative single digit",
			input:    -5,
			expected: false,
		},

		// 零和边界值测试
		{
			name:     "zero",
			input:    0,
			expected: true,
		},
		{
			name:     "max int32 palindrome",
			input:    2147447412,
			expected: true,
		},
		{
			name:     "max int32 non-palindrome",
			input:    2147483647,
			expected: false,
		},

		// 特殊数字测试
		{
			name:     "number ending with zero",
			input:    10,
			expected: false,
		},
		{
			name:     "palindrome with zero middle",
			input:    1002001,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.input)
			if got != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
