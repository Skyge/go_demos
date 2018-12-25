package String

import "strings"

func main()  {
	var s = "A man, a plan, a canal: Panama"
	isPalindrome(s)
}

/**
Given a string, determine if it is a palindrome,
considering only alphanumeric characters and ignoring cases.

Example 1:
	Input: "A man, a plan, a canal: Panama"
	Output: true

Example 2:
	Input: "race a car"
	Output: false

Note: For the purpose of this problem, we define empty string as valid palindrome.
*/
func isAlphaNum(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}

func isCapital(c rune) bool {
	return 'A' <= c && c <= 'Z'
}

func isPalindrome(s string) bool {
	str := ""
	for _, v := range s {
		if isAlphaNum(v) {
			if isCapital(v) {
				str += strings.ToLower(string(v))
			} else {
				str += string(v)
			}
		}
	}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
