package day09_pointer

import (
	"reflect"
	"testing"
)

func TestPointAdd(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		want int
	}{
		{
			name: "nil pointer",
			p:    nil,
			want: 0,
		},
		{
			name: "zero value",
			p:    func() *int { x := 0; return &x }(),
			want: 10,
		},
		{
			name: "positive number",
			p:    func() *int { x := 5; return &x }(),
			want: 15,
		},
		{
			name: "negative number",
			p:    func() *int { x := -3; return &x }(),
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 记录原始值（如果有）
			var originalValue int
			if tt.p != nil {
				originalValue = *tt.p
			}

			got := pointAdd(tt.p)

			// 验证返回值
			if got != tt.want {
				t.Errorf("pointAdd() = %v, want %v", got, tt.want)
			}

			// 验证指针指向的值未被修改（函数不应该修改输入）
			if tt.p != nil && *tt.p != originalValue {
				t.Errorf("pointAdd() modified the input value: got %v, want %v", *tt.p, originalValue)
			}
		})
	}
}

func TestPointMultiply(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty slice",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "single element",
			arr:  []int{3},
			want: []int{6},
		},
		{
			name: "multiple elements",
			arr:  []int{1, 2, 3},
			want: []int{2, 4, 6},
		},
		{
			name: "with zero",
			arr:  []int{0, 5, 10},
			want: []int{0, 10, 20},
		},
		{
			name: "with negative numbers",
			arr:  []int{-1, -2, 3},
			want: []int{-2, -4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入切片的深拷贝
			input := make([]int, len(tt.arr))
			copy(input, tt.arr)

			got := pointMultiply(input)

			// 验证返回值
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pointMultiply() = %v, want %v", got, tt.want)
			}

			// 验证原始输入是否被修改（函数内部可以修改切片元素）
			if !reflect.DeepEqual(input, tt.want) {
				t.Logf("pointMultiply() modifies input slice as expected (got %v)", input)
			}
		})
	}
}
