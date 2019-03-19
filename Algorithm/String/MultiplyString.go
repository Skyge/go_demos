package String

import "bytes"

/**
Given two non-negative integers num1 and num2 represented as strings,
return the product of num1 and num2, also represented as a string.

Example 1:
	Input: num1 = "2", num2 = "3"
	Output: "6"

Example 2:
	Input: num1 = "123", num2 = "456"
	Output: "56088"

Note:
	The length of both num1 and num2 is < 110.
	Both num1 and num2 contain only digits 0-9.
	Both num1 and num2 do not contain any leading zero, except the number 0 itself.
	You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := bytes.NewBufferString("")
	num1Len, num2Len := len(num1), len(num2)
	pos := make([]byte, num1Len+num2Len)
	for i := num1Len - 1; i >= 0; i-- {
		m := num1[i] - '0'
		for j := num2Len - 1; j >= 0; j-- {
			n := num2[j] - '0'
			tmp := m * n
			sum := tmp + pos[i+j+1]
			pos[i+j] += sum / 10
			pos[i+j+1] = sum % 10
		}
	}
	for i := 0; i < len(pos); i++ {
		if i == 0 && pos[i] == 0 {
			continue
		}
		res.WriteByte(pos[i] + '0')
	}
	if res.Len() == 0 {
		return "0"
	}
	return res.String()
}
