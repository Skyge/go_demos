package Other

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
	Input: "()"
	Output: true

Example 2:
	Input: "()[]{}"
	Output: true

Example 3:
	Input: "(]"
	Output: false

Example 4:
	Input: "([)]"
	Output: false

Example 5:
	Input: "{[]}"
	Output: true
*/
func isValid(s string) bool {
	parentheses := map[rune]rune{')': '(', ']': '[', '}': '{'}
	var stack []rune

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if len(stack) > 0 && parentheses[char] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
