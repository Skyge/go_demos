package Array

import (
	"fmt"
)

func main() {
	nums := []int{7,1,5,3,3,6,4}
	output := containsDuplicate(nums)
	fmt.Println(output)
}
/* *
* Given an array of integers,
* find if the array contains any duplicates.
* Your function should return true if any value appears at least twice in the array,
* and it should return false if every element is distinct.
 */
func containsDuplicate(nums []int) bool  {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}