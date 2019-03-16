package String

/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:
	Input: "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3.

Example 2:
	Input: "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.

Example 3:
	Input: "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
				 Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)

	var ret, i, j int
	for j < len(s) {
		if v, ok := m[s[j]]; ok {
			ret = max(ret, len(s[i:j]))
			i = max(i, v+1)
		}
		m[s[j]] = j
		j += 1
	}

	return max(ret, len(s[i:j]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
