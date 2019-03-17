package String

/**
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
In other words, one of the first string's permutations is the substring of the second string.

Example 1:
	Input:s1 = "ab" s2 = "eidbaooo"
	Output:True
	Explanation: s2 contains one permutation of s1 ("ba").

Example 2:
	Input:s1= "ab" s2 = "eidboaoo"
	Output: False

Note:
	The input strings only contain lower case letters.
	The length of both given strings is in range [1, 10,000].
*/
func checkInclusion(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)
	if length1 > length2 {
		return false
	}

	array1 := [128]int{}
	array2 := [128]int{}

	for i := 0; i < length1; i++ {
		array1[s1[i]]++
		array2[s2[i]]++
	}
	if array1 == array2 {
		return true
	}

	for i := length1; i < length2; i++ {
		array2[s2[i]]++
		array2[s2[i-length1]]--
		if array1 == array2 {
			return true
		}
	}

	return false
}
