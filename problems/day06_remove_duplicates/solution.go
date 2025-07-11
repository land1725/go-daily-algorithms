package day06_remove_duplicates

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	m := 1
	for k := 1; k < len(nums); k++ {
		if nums[k] != nums[k-1] {
			nums[m] = nums[k]
			m++
		}
	}
	return m
}
