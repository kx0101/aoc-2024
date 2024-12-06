package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expected := 41
	startX, startY, startDirection := findGuard(grid)
	result := Part1(grid, startX, startY, startDirection)

	if result != expected {
		t.Errorf("Part1 failed: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	grid := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expected := 6
	startX, startY, startDirection := findGuard(grid)
	result := Part2(grid, startX, startY, startDirection)

	if result != expected {
		t.Errorf("Part2 failed: expected %d, got %d", expected, result)
	}
}

func findGuard(grid []string) (int, int, rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '^' || grid[i][j] == 'v' || grid[i][j] == '<' || grid[i][j] == '>' {
				return i, j, rune(grid[i][j])
			}
		}
	}
	return -1, -1, ' '
}
