package Array

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	output := moveZeroes(nums)
	fmt.Println(output)
}

/**
* Given an array nums, write a function to move all 0's to the end of it
* while maintaining the relative order of the non-zero elements.
* Note:
* You must do this in-place without making a copy of the array.
* Minimize the total number of operations.
 */
func moveZeroes(nums []int) []int {
	countZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			countZero++
		} else if countZero != 0 {
			nums[i-countZero] = nums[i]
			nums[i] = 0
		}
	}
	return nums
}
