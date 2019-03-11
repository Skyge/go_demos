package Other

/**
The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.

Note:
	0 ≤ x, y < 231.

Example:
	Input: x = 1, y = 4
	Output: 2
	Explanation:
		1   (0 0 0 1)
		4   (0 1 0 0)
			   ↑   ↑
The above arrows point to positions where the corresponding bits are different.
*/

/**
Thoughts :
z & (z-1) will set all the flipped bits to 0, as 1 & 0 = 0
	00101000
  & 00100111
  = 00100000
therefore we repeat this until z becomes 0, where all set bits are removed.
return the repeated times as number of set bits
*/
func hammingDistance(x int, y int) int {
	result, count := x^y, 0
	for result != 0 {
		result &= result - 1
		count++
	}
	return count
}
