package String

func main() {
	s := "anagram"
	t := "nagaram"
	isAnagram(s, t)
}

/**
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

Example 2:
	Input: s = "rat", t = "car"
	Output: false

Note:
	You may assume the string contains only lowercase alphabets.
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b := make([]int, 26)
	for _, r := range s {
		b[r-'a']++
	}
	for _, r := range t {
		b[r-'a']--
		if b[r-'a'] < 0 {
			return false
		}
	}
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
