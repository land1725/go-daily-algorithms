package day04_longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		// 正常匹配用例
		{
			name:     "示例1-标准前缀",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "完全匹配",
			input:    []string{"prefix", "prefix", "prefix"},
			expected: "prefix",
		},
		{
			name:     "部分匹配",
			input:    []string{"pre", "prefix", "prelude"},
			expected: "pre",
		},

		// 不匹配用例
		{
			name:     "示例2-无公共前缀",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "完全不匹配",
			input:    []string{"abc", "def", "ghi"},
			expected: "",
		},

		// 边界测试
		{
			name:     "空数组",
			input:    []string{},
			expected: "",
		},
		{
			name:     "单元素数组",
			input:    []string{"single"},
			expected: "single",
		},
		{
			name:     "包含空字符串",
			input:    []string{"", "abc", "def"},
			expected: "",
		},
		{
			name:     "长度不同的字符串",
			input:    []string{"ab", "a"},
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := longestCommonPrefix(tt.input)
			assert.Equal(t, tt.expected, actual,
				"输入: %v, 预期: %q, 实际: %q",
				tt.input, tt.expected, actual)
		})
	}
}

// 基准测试
func BenchmarkLongestCommonPrefix(b *testing.B) {
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"prefix", "prefix", "prefix"},
		{"go", "golang", "google"},
	}

	for i := 0; i < b.N; i++ {
		for _, strs := range testCases {
			longestCommonPrefix(strs)
		}
	}
}
