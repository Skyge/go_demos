package String

import "fmt"

func main() {
	var s = "4193 with words"
	fmt.Println(myAtoi(s))
}

/**
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until
the first non-whitespace character is found.
Then, starting from this character,
takes an optional initial plus or minus sign followed by as many numerical digits as possible,
and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number,
which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number,
or if no such sequence exists because either str is empty or it contains only whitespace characters,
no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Example 1:
	Input: "42"
	Output: 42

Example 2:
	Input: "   -42"
	Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
Then take as many numerical digits as possible, which gets 42.

Example 3:
	Input: "4193 with words"
	Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

Example 4:
	Input: "words and 987"
	Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
digit or a +/- sign. Therefore no valid conversion could be performed.

Example 5:
	Input: "-91283472332"
	Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
Therefore INT_MIN (−231) is returned.

Note:
	Only the space character ' ' is considered as whitespace character.
	Assume we are dealing with an environment
	which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
	If the numerical value is out of the range of representable values,
	INT_MAX (231 − 1) or INT_MIN (−231) is returned.
*/
const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func myAtoi(str string) int {
	var out int64
	sign, isParsing := "", false

	for _, s := range str {
		if s == ' ' {
			if isParsing {
				break
			}
			continue
		}
		if s == '+' || s == '-' {
			if isParsing || sign != "" {
				break
			}
			isParsing = true
			sign = string(s)
			continue
		}
		if !(s >= 48 && s <= 57) {
			break
		}

		isParsing = true
		out *= 10
		out += int64(s - 48)
		if out >= maxInt32 || -out <= minInt32 {
			break
		}
	}

	if sign == "-" {
		out = -out
		if out < minInt32 {
			return minInt32
		}
	}
	if out > maxInt32 {
		return maxInt32
	}
	return int(out)
}
