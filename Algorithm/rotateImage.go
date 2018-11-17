package main

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
}

/*
* You are given an n x n 2D matrix representing an image.
* Rotate the image by 90 degrees (clockwise).
* Note:
* You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
* DO NOT allocate another 2D matrix and do the rotation.
 */
func rotate(matrix [][]int) {
	n := len(matrix)

	// transpose
	for x := 0; x < n; x++ {
		for y := 0; y < x; y++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}

	}

	// horizontal flip
	for y := 0; y < n; y++ {
		for x := 0; x < n/2; x++ {
			matrix[y][x], matrix[y][n-1-x] = matrix[y][n-1-x], matrix[y][x]
		}
	}
}
