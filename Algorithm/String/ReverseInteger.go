package String

import (
	"strconv"
)

func main() {
	input := -1234567
	reverseInteger(input)
}

/**Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
	Input: 123
	Output: 321

Example 2:
	Input: -123
	Output: -321

Example 3:
	Input: 120
	Output: 21
Note:
	Assume we are dealing with an environment which could only store integers
	within the 32-bit signed integer range: [−231,  231 − 1].
	For the purpose of this problem,
	assume that your function returns 0 when the reversed integer overflows.
*/

func reverseInteger(x int) int {
	var result = ""
	if x < 0 {
		result = "-"
		x = -x
	}
	for x/10 != 0 {
		result = result + strconv.Itoa(x%10)
		x = x / 10
	}
	result = result + strconv.Itoa(x%10)
	a, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		return 0
	}

	return int(a)
}
