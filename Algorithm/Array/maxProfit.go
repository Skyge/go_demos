package Array

import "fmt"

func main() {
	nums := []int{7, 1, 5, 6, 3, 6, 4}
	output := maxProfit(nums)
	fmt.Println(output)
}

/**
 * Say you have an array for which the ith element is the price of a given stock on day i.
 * Design an algorithm to find the maximum profit.
 * You may complete as many transactions as you like：
 * (i.e., buy one and sell one share of the stock multiple times).
 * Note: You may not engage in multiple transactions at the same time：
 * (i.e., you must sell the stock before you buy again).
 */
func maxProfit(nums []int) int {
	lens := len(nums)
	profit := 0
	for i := 0; i < lens-1; i++ {
		if nums[i] < nums[i+1] {
			profit += nums[i+1] - nums[i]
		}
	}
	return profit
}
