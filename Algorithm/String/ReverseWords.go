package String

/**
Given an input string, reverse the string word by word.

Example 1:
	Input: "the sky is blue"
	Output: "blue is sky the"

Example 2:
	Input: "  hello world!  "
	Output: "world! hello"
	Explanation: Your reversed string should not contain leading or trailing spaces.

Example 3:
	Input: "a good   example"
	Output: "example good a"
	Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

Note:
	A word is defined as a sequence of non-space characters.
	Input string may contain leading or trailing spaces.
	However, your reversed string should not contain leading or trailing spaces.
	You need to reduce multiple spaces between two words to a single space in the reversed string.
*/
func reverseWords(s string) string {
	a := []rune(s)
	b := make([]rune, 1)
	temp := 0

	for i := 0; i < len(a); i++ {
		if len(b) == 1 && a[i] == ' ' {
			continue
		} else if len(b) == 1 && a[i] != ' ' {
			b = append(b, a[i])
		} else if a[i] != ' ' {
			if temp != 0 {
				b = append(b, ' ')
				temp = 0
			}
			b = append(b, a[i])
		} else if a[i] == ' ' {
			temp += 1
			continue
		}
	}

	res := make([]string, 1)
	tempStr := ""
	for i := 1; i < len(b); i++ {
		if b[i] == ' ' {
			res = append(res, tempStr)
			res = append(res, " ")
			tempStr = ""
			continue
		}
		tempStr += string(b[i])
	}
	res = append(res, tempStr)

	result := ""
	for i := len(res) - 1; i > 0; i-- {
		result += res[i]
	}

	return result
}
