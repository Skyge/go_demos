package Array

import "fmt"

func main() {
	nums := []int{-7, 7, 7, 3, 3, 4, 4}
	output := singleNumber(nums)
	fmt.Println(output)
}

/* *
* Given a non-empty array of integers,
* every element appears twice except for one. Find that single one.
*
* Note:Your algorithm should have a linear runtime complexity.
* Could you implement it without using extra memory?
 */
func singleNumber(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num ^= nums[i]
	}
	return num
}
