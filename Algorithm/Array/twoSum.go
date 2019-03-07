package Array

import (
	f "fmt"
)

func main() {
	nums := []int{2, 8, 11, 15, 7}
	target := 9
	output := twoSum(nums, target)
	f.Println(output)
}

/**
 * Given an array of integers,
 * return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution,
 * and you may not use the same element twice.
 */
func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	temp := make(map[int]int)
	for i, num := range nums {
		if expected, exited := temp[target-num]; exited {
			result[0] = expected
			result[1] = i
			return result
		}
		temp[num] = i
	}
	return result
}
