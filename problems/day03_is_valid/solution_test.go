package day03_is_valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// 常规有效用例
		{
			name:     "简单括号匹配",
			input:    "()",
			expected: true,
		},
		{
			name:     "混合括号匹配",
			input:    "()[]{}",
			expected: true,
		},
		{
			name:     "嵌套括号",
			input:    "([{}])",
			expected: true,
		},

		// 常规无效用例
		{
			name:     "不匹配的括号",
			input:    "(]",
			expected: false,
		},
		{
			name:     "顺序错误",
			input:    "([)]",
			expected: false,
		},
		{
			name:     "只有闭括号",
			input:    "])}",
			expected: false,
		},

		// 边界测试
		{
			name:     "空字符串",
			input:    "",
			expected: true, // 根据业务逻辑决定 true/false
		},
		{
			name:     "单字符-开括号",
			input:    "(",
			expected: false,
		},
		{
			name:     "单字符-闭括号",
			input:    "]",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 每次重置栈状态
			actual := isValid(tt.input)
			assert.Equal(t, tt.expected, actual,
				"输入: %q, 预期: %v, 实际: %v",
				tt.input, tt.expected, actual)
		})
	}
}

// 基准测试
func BenchmarkIsValid(b *testing.B) {
	testCases := []string{
		"()[]{}",
		"([{}])",
		"([{}{}()][]([]))",
	}

	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			isValid(s)
		}
	}
}
