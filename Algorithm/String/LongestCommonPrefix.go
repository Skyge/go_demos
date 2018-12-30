package String

import "fmt"

func main()  {
	input := []string{
		"flower",
		"flow",
		"flight",
	}

	fmt.Println(longestCommonPrefix(input))
}

/**
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
	Input: ["flower","flow","flight"]
	Output: "fl"

Example 2:
	Input: ["dog","racecar","car"]
	Output: ""
Explanation: There is no common prefix among the input strings.

Note:
	All given inputs are in lowercase letters a-z.
 */
func longestCommonPrefix(strs []string) (result string) {
	shortest := -1
	for _, s := range strs {
		if len(s) < shortest || shortest == -1{
			shortest = len(s)
		}
	}
	for i := 0; i < shortest; i++ {
		for _, s := range strs {
			if s[i] != strs[0][i] {
				return
			}
		}
		result += string(strs[0][i])
	}
	return
}