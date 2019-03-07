package DynamicProgramming

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which
has the largest sum and return its sum.

Example:
	Input: [-2,1,-3,4,-1,2,1,-5,4],
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:
	If you have figured out the O(n) solution, t
	ry coding another solution using the divide and conquer approach, which is more subtle.
*/
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := max

	for _, n := range nums[1:] {
		if sum+n > n {
			sum += n
		} else {
			sum = n
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
