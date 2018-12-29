package String

import "fmt"

func main()  {
	input := 5
	fmt.Println(countAndSay(input))
}

/**
The count-and-say sequence is the sequence of integers with the first five terms as following:
	1.     1
	2.     11
	3.     21
	4.     1211
	5.     111221

	1 is read off as "one 1" or 11.
	11 is read off as "two 1s" or 21.
	21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30,
generate the nth term of the count-and-say sequence.

Note:
	Each term of the sequence of integers will be represented as a string.

Example 1:
	Input: 1
	Output: "1"

Example 2:
	Input: 4
	Output: "1211"
 */
func countAndSay(n int) string {
	return gen("1", n-1)
}

func gen(str string, n int) string {
	if n == 0 {
		return str
	}

	var result []byte
	index := 0
	for index < len(str) {
		c := str[index]
		last := index
		for j := index; j < len(str); j++ {
			if str[j] != c {
				break
			}
			last = j
		}

		n := last - index + 1
		result = append(result, byte(n) + '0')
		result = append(result, c)
		index = last + 1
	}

	return gen(string(result), n - 1)
}