package Array

/**
Given a non-empty 2D array grid of 0's and 1's,
an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:
	[[0,0,1,0,0,0,0,1,0,0,0,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,1,1,0,1,0,0,0,0,0,0,0,0],
	 [0,1,0,0,1,1,0,0,1,0,1,0,0],
	 [0,1,0,0,1,1,0,0,1,1,1,0,0],
	 [0,0,0,0,0,0,0,0,0,0,1,0,0],
	 [0,0,0,0,0,0,0,1,1,1,0,0,0],
	 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
	Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.

Example 2:
	[[0,0,0,0,0,0,0,0]]
	Given the above grid, return 0.

Note: The length of each dimension in the given grid does not exceed 50.
*/
func maxAreaOfIsland(grid [][]int) int {
	max := 0

	for i, row := range grid {
		for j, col := range row {
			if col == 0 {
				continue
			}

			area := areaOfIsland(grid, i, j)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func areaOfIsland(grid [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) ||
		grid[i][j] == 0 {
		return 0
	}

	grid[i][j] = 0

	total := 1

	total += areaOfIsland(grid, i-1, j) // up
	total += areaOfIsland(grid, i+1, j) // down
	total += areaOfIsland(grid, i, j-1) // left
	total += areaOfIsland(grid, i, j+1) // right

	return total
}
