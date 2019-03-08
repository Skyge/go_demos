package Math

/**
Given an integer, write a function to determine if it is a power of three.

Example 1:
	Input: 27
	Output: true
Example 2:
	Input: 0
	Output: false
Example 3:
	Input: 9
	Output: true
Example 4:
	Input: 45
	Output: false

Follow up:
	Could you do it without using any loop / recursion?
*/
//func isPowerOfThree(n int) bool {
//	if n < 1 {
//		return false
//	} else if n == 1 {
//		return true
//	}
//	for n != 1 {
//		if n%3 != 0 {
//			return false
//		}
//		n /= 3
//	}
//	return true
//}
// do it without using any loop / recursion
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0  // 1162261467 = 2^31 - 1
}
