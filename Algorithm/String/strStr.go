package String

import "fmt"

func main() {
	var haystack = "aaaaa"
	var needle = "bba"
	fmt.Println(strStr(haystack, needle))
}

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack,
or -1 if needle is not part of haystack.

Example 1:
	Input: haystack = "hello", needle = "ll"
	Output: 2

Example 2:
	Input: haystack = "aaaaa", needle = "bba"
	Output: -1

Clarification:
	What should we return when needle is an empty string?
	This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string.
This is consistent to C's strstr() and Java's indexOf().
*/
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	searchOffset := len(needle)
	if searchOffset > haystackLen {
		return -1
	}
	for i := 0; i <= haystackLen; i++ {
		if i+searchOffset > haystackLen {
			break
		}
		if haystack[i:i+searchOffset] == needle {
			return i
		}
	}

	return -1
}
