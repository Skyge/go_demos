package Other

/**
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
	Input: 5
	Output:
	[
		 [1],
		[1,1],
	   [1,2,1],
	  [1,3,3,1],
	 [1,4,6,4,1]
	]
*/
func generate(numRows int) [][]int {
	s := make([][]int, numRows)
	if numRows == 0 {
		return s
	}
	for i := 1; i <= numRows; i++ {
		s[i-1] = make([]int, i)
		s[i-1][0] = 1
		s[i-1][i-1] = 1
		if i-1 > 1 {
			for j := 1; j < i-1; j++ {
				s[i-1][j] = s[i-2][j-1] + s[i-2][j]
			}
		}
	}
	return s
}
