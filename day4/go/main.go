package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dx = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	dy = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
)

func main() {
	grid := ReadGridFromFile("../input.txt")

	count := CountXMASOccurancies(grid)
	count2 := CountXMASPatterns(grid)

	fmt.Printf("part 1: %d\n", count)
	fmt.Printf("part 2: %d\n", count2)
}

func ReadGridFromFile(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file: ", err)
		return nil
	}

	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file: ", err)
		return nil
	}

	return grid
}

func CountXMASOccurancies(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	needle := "XMAS"
	count := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for direction := 0; direction < 8; direction++ {
				if checkWord(grid, i, j, direction, needle) {
					count++
				}
			}
		}
	}

	return count
}

func checkWord(grid [][]rune, row, col, direction int, needle string) bool {
	rows := len(grid)
	cols := len(grid[0])

	for k := 0; k < len(needle); k++ {
		newRow := row + k*dx[direction]
		newCol := col + k*dy[direction]

		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || grid[newRow][newCol] != rune(needle[k]) {
			return false
		}

	}

	return true
}

func CountXMASPatterns(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if checkBothDiagonals(grid, i, j) {
				count++
			}
		}
	}

	return count
}

func checkBothDiagonals(grid [][]rune, i, j int) bool {
	// first diagonal \
	// top left, middle, bottom right
	diag1MAS := IsMASSequence(grid[i-1][j-1], grid[i][j], grid[i+1][j+1])
	diag1SAM := IsMASSequence(grid[i+1][j+1], grid[i][j], grid[i-1][j-1])

	// second diagonal /
	// top right, middle, bottom left
	diag2MAS := IsMASSequence(grid[i-1][j+1], grid[i][j], grid[i+1][j-1])
	diag2SAM := IsMASSequence(grid[i+1][j-1], grid[i][j], grid[i-1][j+1])

	return (diag1MAS || diag1SAM) && (diag2MAS || diag2SAM)
}

func IsMASSequence(m, a, s rune) bool {
	return m == 'M' && a == 'A' && s == 'S'
}
