package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	guardX, guardY := -1, -1
	guardDirection := ' '

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {

			if grid[i][j] == '^' || grid[i][j] == 'v' || grid[i][j] == '<' || grid[i][j] == '>' {
				guardX, guardY = i, j
				guardDirection = rune(grid[i][j])
				grid[i] = grid[i][:j] + "." + grid[i][j+1:]

				break
			}
		}
	}

	if guardX == -1 || guardY == -1 {
		fmt.Println("Guard's starting position not found.")
		return
	}

	part1Result := Part1(grid, guardX, guardY, guardDirection)
	part2Result := Part2(grid, guardX, guardY, guardDirection)

	fmt.Printf("Part 1: %d\n", part1Result)
	fmt.Printf("Part 2: %d\n", part2Result)
}

func Part1(grid []string, startX, startY int, startDirection rune) int {
	directions := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	turnRight := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true

	x, y := startX, startY
	direction := startDirection

	rows := len(grid)
	cols := len(grid[0])

	for {
		dx, dy := directions[direction][0], directions[direction][1]
		nextX, nextY := x+dx, y+dy

		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			break
		}

		if grid[nextX][nextY] == '#' {
			direction = turnRight[direction]
		} else {
			x, y = nextX, nextY
			visited[[2]int{x, y}] = true
		}
	}

	return len(visited)
}

// gpt solution i couldnt figure it out smh
func Part2(grid []string, startX, startY int, startDirection rune) int {
	directions := map[rune][2]int{
		'^': {-1, 0},
		'v': {1, 0},
		'<': {0, -1},
		'>': {0, 1},
	}

	turnRight := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	rows := len(grid)
	cols := len(grid[0])
	loopObstructionCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != '.' || (i == startX && j == startY) {
				continue
			}

			if simulateWithObstruction(grid, startX, startY, startDirection, i, j, directions, turnRight) {
				loopObstructionCount++
			}
		}
	}

	return loopObstructionCount
}

func simulateWithObstruction(grid []string, startX, startY int, startDirection rune, obstructionX, obstructionY int,
	directions map[rune][2]int, turnRight map[rune]rune) bool {

	gridClone := make([]string, len(grid))
	for i := range grid {
		gridClone[i] = grid[i]
	}

	gridClone[obstructionX] = gridClone[obstructionX][:obstructionY] + "#" + gridClone[obstructionX][obstructionY+1:]

	visitedStates := make(map[[3]interface{}]bool)
	visitedStates[[3]interface{}{startX, startY, startDirection}] = true

	x, y := startX, startY
	direction := startDirection

	rows := len(gridClone)
	cols := len(gridClone[0])

	for {
		dx, dy := directions[direction][0], directions[direction][1]
		nextX, nextY := x+dx, y+dy

		if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
			return false
		}

		if gridClone[nextX][nextY] == '#' {
			direction = turnRight[direction]
		} else {
			x, y = nextX, nextY

			state := [3]interface{}{x, y, direction}
			if visitedStates[state] {
				return true
			}

			visitedStates[state] = true
		}
	}
}
