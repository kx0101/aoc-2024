package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, rows, cols := ReadInput("../input.txt")

	part1Total := Part1(grid, rows, cols)
	fmt.Printf("Part 1 total score: %d\n", part1Total)

	part2Total := Part2(grid, rows, cols)
	fmt.Printf("Part 2 total score: %d\n", part2Total)
}

func Part1(grid [][]int, rows, cols int) int {
	totalScore := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				visited := make([][]bool, rows)

				for k := range visited {
					visited[k] = make([]bool, cols)
				}

				totalScore += GetTrailheadScore(grid, i, j, rows, cols, visited)
			}
		}
	}

	return totalScore
}

func Part2(grid [][]int, rows, cols int) int {
	totalRating := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				memo := make(map[[2]int]int)

				totalRating += GetTrailheadRating(grid, i, j, rows, cols, 0, memo)
			}
		}
	}

	return totalRating
}

func GetTrailheadRating(grid [][]int, row, col, rows, cols int, currHeight int, memo map[[2]int]int) int {
	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] > 9 || grid[row][col] != currHeight {
		return 0
	}

	if grid[row][col] == 9 {
		return 1
	}

	key := [2]int{row, col}
	if value, exists := memo[key]; exists {
		return value
	}

	totalTrails := 0

	totalTrails += GetTrailheadRating(grid, row-1, col, rows, cols, currHeight+1, memo)
	totalTrails += GetTrailheadRating(grid, row+1, col, rows, cols, currHeight+1, memo)
	totalTrails += GetTrailheadRating(grid, row, col-1, rows, cols, currHeight+1, memo)
	totalTrails += GetTrailheadRating(grid, row, col+1, rows, cols, currHeight+1, memo)

	memo[key] = totalTrails

	return totalTrails
}

func GetTrailheadScore(grid [][]int, row, col, rows, cols int, visited [][]bool) int {
	reachableNines := make(map[[2]int]struct{})

	dfs(grid, row, col, rows, cols, visited, 0, reachableNines)

	return len(reachableNines)
}

func dfs(grid [][]int, row, col, rows, cols int, visited [][]bool, currentHeight int, reachableNines map[[2]int]struct{}) {
	if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] || grid[row][col] != currentHeight {
		return
	}

	visited[row][col] = true

	if grid[row][col] == 9 {
		reachableNines[[2]int{row, col}] = struct{}{}
		return
	}

	dfs(grid, row-1, col, rows, cols, visited, currentHeight+1, reachableNines)
	dfs(grid, row+1, col, rows, cols, visited, currentHeight+1, reachableNines)
	dfs(grid, row, col-1, rows, cols, visited, currentHeight+1, reachableNines)
	dfs(grid, row, col+1, rows, cols, visited, currentHeight+1, reachableNines)
}

func ReadInput(filename string) ([][]int, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for i, c := range line {
			row[i] = int(c - '0')
		}

		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rows := len(grid)
	cols := len(grid[0])

	return grid, rows, cols
}
