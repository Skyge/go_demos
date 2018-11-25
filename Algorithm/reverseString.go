package main

func main() {
	s := "Hello World"
	reverseString(s)
}

/**
Write a function that takes a string as input and returns the string reversed.

Example 1:
	Input: "hello"
	Output: "olleh"

Example 2:
	Input: "A man, a plan, a canal: Panama"
	Output: "amanaP :lanac a ,nalp a ,nam A"
 */

func reverseString(s string) string {
	a := []rune(s)
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	return string(a)
}
