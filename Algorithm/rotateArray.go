package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	output := rotate(nums, k)
	fmt.Println(output)
}

/**
* Given an array, rotate the array to the right by k steps, where k is non-negative.
 */
func rotate(nums []int, k int) []int {
	k = k % len(nums)
	if k != 0 {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
	}
	return nums
}
