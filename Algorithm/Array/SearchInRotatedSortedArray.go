package Array

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

Example 2:
	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
*/
func binarySearch(nums []int, target int, left int, right int) int {
	n := right - left + 1

	if n == 0 {
		return -1
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	if left >= right {
		return -1
	}

	mid := (right + left) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[left] < target && nums[mid] > target {
		return binarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < target && nums[right] > target {
		return binarySearch(nums, target, mid+1, right)
	}

	if nums[mid] > nums[right] {
		return binarySearch(nums, target, mid+1, right)
	}
	return binarySearch(nums, target, left, mid-1)
}

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}
