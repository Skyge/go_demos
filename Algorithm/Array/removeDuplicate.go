package Array

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	numsLen := removeDuplicates(nums)
	fmt.Println(numsLen)
	fmt.Println(nums)
}

/**
 * Given a sorted array nums,
 * remove the duplicates in-place such that each element appear only once and return the new length.
 * Do not allocate extra space for another array,
 * you must do this by modifying the input array in-place with O(1) extra memory.
 */
func removeDuplicates(nums []int) int {
	number := 0
	lens := len(nums)
	for i := 0; i < lens; i++ {
		if nums[i] != nums[number] {
			number++
			nums[number] = nums[i]
		}
	}
	return number + 1
}
