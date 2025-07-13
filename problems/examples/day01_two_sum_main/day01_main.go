package main

import (
	"fmt"
	"github.com/land1725/go-daily-algorithms/problems/day01_two_sum"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := day01_two_sum.TwoSum(nums, target)
	fmt.Printf("Input: %v\nTarget: %d\nResult: %v\n", nums, target, result)
}
