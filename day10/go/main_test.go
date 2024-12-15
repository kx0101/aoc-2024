package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "first",
			grid: [][]int{
				{0, 1, 2, 3},
				{1, 2, 3, 4},
				{8, 7, 6, 5},
				{9, 8, 7, 6},
			},
			expected: 1,
		},
		{
			name: "second",
			grid: [][]int{
				{0, 0, 0},
				{1, 2, 1},
				{2, 3, 2},
				{9, 9, 9},
			},
			expected: 0,
		},
		{
			name: "third",
			grid: [][]int{
				{8, 9, 0, 1, 0, 1, 2, 3},
				{7, 8, 1, 2, 1, 8, 7, 4},
				{8, 7, 4, 3, 0, 9, 6, 5},
				{9, 6, 5, 4, 9, 8, 7, 4},
				{4, 5, 6, 7, 8, 9, 0, 3},
				{3, 2, 0, 1, 9, 0, 1, 2},
				{0, 1, 3, 2, 9, 8, 0, 1},
				{1, 0, 4, 5, 6, 7, 3, 2},
			},
			expected: 36,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			rows, cols := len(test.grid), len(test.grid[0])
			result := Part1(test.grid, rows, cols)

			if result != test.expected {
				t.Errorf("got=%d, want=%d", result, test.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "first",
			grid: [][]int{
				{0, 0, 0},
				{1, 2, 1},
				{2, 3, 2},
				{9, 9, 9},
			},
			expected: 0,
		},
		{
			name: "second",
			grid: [][]int{
				{8, 9, 0, 1, 0, 1, 2, 3},
				{7, 8, 1, 2, 1, 8, 7, 4},
				{8, 7, 4, 3, 0, 9, 6, 5},
				{9, 6, 5, 4, 9, 8, 7, 4},
				{4, 5, 6, 7, 8, 9, 0, 3},
				{3, 2, 0, 1, 9, 0, 1, 2},
				{0, 1, 3, 2, 9, 8, 0, 1},
				{1, 0, 4, 5, 6, 7, 3, 2},
			},
			expected: 81,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rows, cols := len(test.grid), len(test.grid[0])
			result := Part2(test.grid, rows, cols)

			if result != test.expected {
				t.Errorf("got=%d, want=%d", result, test.expected)
			}
		})
	}
}
