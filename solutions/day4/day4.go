package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadGrid(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, line)
		}
	}
	return grid, scanner.Err()
}

func copyGrid(grid []string) [][]byte {
	out := make([][]byte, len(grid))
	for i := range grid {
		out[i] = []byte(grid[i])
	}
	return out
}

func countAdj(grid [][]byte, r, c int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, d := range directions {
		nr, nc := r+d[0], c+d[1]
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			if grid[nr][nc] == '@' {
				count++
			}
		}
	}
	return count
}

func Part1(grid []string) int {
	byteGrid := copyGrid(grid)
	accessible := 0

	for r, row := range byteGrid {
		for c, cell := range row {
			if cell == '@' && countAdj(byteGrid, r, c) < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func Part2(grid []string) int {
	byteGrid := copyGrid(grid)
	totalRemoved := 0

	for {
		var toRemove [][2]int

		for r, row := range byteGrid {
			for c, cell := range row {
				if cell == '@' && countAdj(byteGrid, r, c) < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		for _, cell := range toRemove {
			r, c := cell[0], cell[1]
			byteGrid[r][c] = '.'
		}

		totalRemoved += len(toRemove)
	}

	return totalRemoved
}

func main() {
	grid, err := loadGrid("solutions/day4/day4.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Part 1:", Part1(grid))
	fmt.Println("Part 2:", Part2(grid))
}
