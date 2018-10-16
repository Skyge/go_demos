package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	//output := rotate(nums, k)
	//fmt.Println(output)
	rotate2(nums, k)
	fmt.Println(nums)
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

// method2
func rotate2(nums []int, k int) {
	// 假定 k >= 0

	n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
