package Math

/**
Count the number of prime numbers less than a non-negative number, n.

Example:
	Input: 10
	Output: 4
	Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	dict := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !dict[i] {
			count++
			for j := 2; i*j < n; j++ {
				dict[i*j] = true
			}
		}
	}
	return count
}
