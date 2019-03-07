package Array

import "fmt"

func main() {
	nums := []int{9, 9, 9}
	output := plusOne(nums)
	fmt.Println(output)
}

/**
 * Given a non-empty array of digits representing a non-negative integer,
 * plus one to the integer.
 * The digits are stored such that the most significant digit is at the head of the list,
 * and each element in the array contain a single digit.
 * You may assume the integer does not contain any leading zero,
 * except the number 0 itself.
 */
func plusOne(nums []int) []int {
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i] += 1
			return nums
		} else {
			nums[i] = 0
		}
	}
	result := make([]int, n+1)
	result[0] = 1
	return result
}
