package main

import (
	"fmt"
	"strconv"
)

func main() {
	board := [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}

	fmt.Println(isValidSudoku(board))
}

/*
* Determine if a 9x9 Sudoku board is valid.
* Only the filled cells need to be validated according to the following rules:
*	Each row must contain the digits 1-9 without repetition.
*	Each column must contain the digits 1-9 without repetition.
*	Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.
* A partially filled sudoku which is valid.
* The Sudoku board could be partially filled, where empty cells are filled with the character '.'.
 */
func isValidSudoku(board [][]byte) bool {
	// row
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// column
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// 3X3 cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for row := i * 3; row < i*3+3; row++ {
				for column := j * 3; column < j*3+3; column++ {
					cellVal := board[row][column]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}
	return true
}
