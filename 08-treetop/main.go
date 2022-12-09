package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid [][]int

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		row := make([]int, 0, len(s.Text()))
		for _, r := range s.Text() {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}

	bestScore := 0
	for row := range grid {
		for col := range grid[row] {
			score := scenicScore(grid, row, col)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println(bestScore)
}

func scenicScore(grid [][]int, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	treesize := grid[row][col]

	if row == 0 || row == rows-1 || col == 0 || col == cols-1 {
		return 0
	}

	if row == 3 && col == 2 {
		_ = row
	}

	var (
		up, down, left, right int
	)

	for r := row - 1; r >= 0; r-- {
		up++
		if grid[r][col] >= treesize {
			break
		}
	}

	for r := row + 1; r < rows; r++ {
		down++
		if grid[r][col] >= treesize {
			break
		}
	}

	for c := col - 1; c >= 0; c-- {
		left++
		if grid[row][c] >= treesize {
			break
		}
	}

	for c := col + 1; c < cols; c++ {
		right++
		if grid[row][c] >= treesize {
			break
		}
	}

	return left * right * up * down

}
