package MergeSortedArray

/**
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
	The number of elements initialized in nums1 and nums2 are m and n respectively.
	You may assume that nums1 has enough space (size that is greater or equal to m + n)
	to hold additional elements from nums2.

Example:
	Input:
		nums1 = [1,2,3,0,0,0], m = 3
		nums2 = [2,5,6],       n = 3
	Output: [1,2,2,3,5,6]
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1
	i1 := m - 1
	i2 := n - 1
	for i1 >= 0 && i2 >= 0 {
		n1 := nums1[i1]
		n2 := nums2[i2]
		if n2 > n1 {
			nums1[idx] = n2
			i2--
		} else {
			nums1[idx] = n1
			i1--
		}
		idx--
	}

	for i2 >= 0 {
		nums1[idx] = nums2[i2]
		idx--
		i2--
	}
}
