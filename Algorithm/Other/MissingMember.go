package Other

/**
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n,
find the one that is missing from the array.

Example 1:
	Input: [3,0,1]
	Output: 2

Example 2:
	Input: [9,6,4,2,3,5,7,0,1]
	Output: 8
Note:
Your algorithm should run in linear runtime complexity.
Could you implement it using only constant extra space complexity?
*/
// method1:
//func missingNumber(nums []int) (ret int) {
//	n := len(nums)
//	if n == 0 {
//		return
//	}
//
//	sum := 0
//	for _, v := range nums {
//		sum += v
//	}
//
//	return (n+1)*n/2 - sum // (0+n)*(n+1)/2  = (n+1)*n/2
//}
// method2
func missingNumber(nums []int) int {
	// n^m^n = m  so (0^1^2^3...^n)  ^ (0^1^3...^n) = 2(missing number)
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
		res ^= i
	}

	return res
}
