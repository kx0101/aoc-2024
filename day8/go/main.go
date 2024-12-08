package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	input := readInput("../input.txt")
	rows := len(input)
	cols := len(input[0])

	antennas := make(map[rune][][2]int)
	for y, line := range input {
		for x, char := range line {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				antennas[char] = append(antennas[char], [2]int{x, y})
			}
		}
	}

	antinodes := make(map[[2]int]bool)
	for _, positions := range antennas {
		for _, pos := range positions {
			antinodes[pos] = true
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				x1, y1 := positions[i][0], positions[i][1]
				x2, y2 := positions[j][0], positions[j][1]

				addAntinodes(x1, y1, x2, y2, rows, cols, antinodes)
			}
		}
	}

	fmt.Println(len(antinodes))
}

func addAntinodes(x1, y1, x2, y2, rows, cols int, antinodes map[[2]int]bool) {
	if x1 == x2 {
		minY := min(y1, y2)
		maxY := max(y1, y2)

		for y := minY; y <= maxY; y++ {
			if isWithinBounds(x1, y, rows, cols) {
				antinodes[[2]int{x1, y}] = true
			}
		}
	} else if y1 == y2 {
		minX := min(x1, x2)
		maxX := max(x1, x2)

		for x := minX; x <= maxX; x++ {
			if isWithinBounds(x, y1, rows, cols) {
				antinodes[[2]int{x, y1}] = true
			}
		}
	} else {
		dx := x2 - x1
		dy := y2 - y1

		// traverse from (x1, y1) to (x2, y2)
		traverseDiagonal(x1, y1, dx, dy, rows, cols, antinodes)

		// traverse from (x, y2) to (x1, y1)
		traverseDiagonal(x2, y2, -dx, -dy, rows, cols, antinodes)
	}
}

func traverseDiagonal(x, y, dx, dy, rows, cols int, antinodes map[[2]int]bool) {
	for {
		if !isWithinBounds(x, y, rows, cols) {
			break
		}

		antinodes[[2]int{x, y}] = true

		x += dx
		y += dy
	}
}

func isWithinBounds(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < cols && y < rows
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return input
}
