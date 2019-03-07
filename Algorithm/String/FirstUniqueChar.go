package String

import "fmt"

func main() {
	str := "abcddc"
	firstUniqChar := firstUniqChar(str)
	fmt.Println("The first non-repeating character is:", firstUniqChar)
}

/**
* Given a string, find the first non-repeating character in it and return it's index.
* If it doesn't exist, return -1.
*
* Examples:
*	s = "leetcode"
*	return 0.
*
*	s = "loveleetcode",
*	return 2.

*	Note: You may assume the string contain only lowercase letters.
 */

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	arr := make([]int, 26)
	for _, v := range s {
		arr[byte(v)-'a']++
	}

	for i, v := range s {
		if arr[byte(v)-'a'] == 1 {
			return i
		}
	}
	return -1
}
