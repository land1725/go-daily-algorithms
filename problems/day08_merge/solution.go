package day08_merge

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按区间起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		if last[1] < current[0] {
			// 无重叠，直接添加
			merged = append(merged, current)
		} else {
			// 合并区间
			if current[1] > last[1] {
				last[1] = current[1]
			}
		}
	}

	return merged
}
