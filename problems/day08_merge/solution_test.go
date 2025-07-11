package day08_merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "empty input",
			intervals: [][]int{},
			want:      [][]int{},
		},
		{
			name:      "single interval",
			intervals: [][]int{{1, 3}},
			want:      [][]int{{1, 3}},
		},
		{
			name:      "non-overlapping intervals",
			intervals: [][]int{{1, 3}, {4, 6}, {8, 10}},
			want:      [][]int{{1, 3}, {4, 6}, {8, 10}},
		},
		{
			name:      "completely overlapping intervals",
			intervals: [][]int{{1, 4}, {2, 3}},
			want:      [][]int{{1, 4}},
		},
		{
			name:      "partially overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "adjacent intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "multiple merges needed",
			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			want:      [][]int{{1, 10}},
		},
		{
			name:      "unsorted intervals",
			intervals: [][]int{{8, 10}, {15, 18}, {1, 3}, {2, 6}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the original
			input := make([][]int, len(tt.intervals))
			copy(input, tt.intervals)

			got := merge(input)

			// Verify the result
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}

			//// Verify the original input wasn't modified
			//if !reflect.DeepEqual(input, tt.intervals) {
			//	t.Errorf("merge() modified the input: got %v, want %v", input, tt.intervals)
			//}
		})
	}
}

// Helper function to check if intervals are sorted and non-overlapping
func TestMergeResultProperties(t *testing.T) {
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {0, 2}, {3, 5}},
	}

	for _, intervals := range testCases {
		t.Run("", func(t *testing.T) {
			result := merge(intervals)

			// 1. Check if intervals are sorted
			for i := 1; i < len(result); i++ {
				if result[i-1][0] > result[i][0] {
					t.Errorf("intervals not sorted: %v", result)
				}
			}

			// 2. Check no overlapping intervals exist
			for i := 1; i < len(result); i++ {
				if result[i-1][1] >= result[i][0] {
					t.Errorf("overlapping intervals found: %v and %v", result[i-1], result[i])
				}
			}
		})
	}
}
