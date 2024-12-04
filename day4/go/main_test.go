package main

import "testing"

func TestCountXMASOccurancies(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]rune
		expected int
	}{
		{
			name: "single xmas",
			grid: [][]rune{
				{'X', 'M', 'A', 'S'},
			},
			expected: 1,
		},
		{
			name: "multiple xmas in a row",
			grid: [][]rune{
				{'X', 'M', 'A', 'S', 'X', 'M', 'A', 'S'},
			},
			expected: 2,
		},
		{
			name: "no xmas",
			grid: [][]rune{
				{'X', 'A', 'S', 'M', 'A', 'S'},
			},
			expected: 0,
		},
		{
			name: "full grid with xmas",
			grid: [][]rune{
				{'M', 'A', 'S', 'X', 'M', 'A', 'S'},
				{'A', 'M', 'S', 'X', 'M', 'A', 'S'},
				{'S', 'X', 'M', 'A', 'S', 'X', 'M'},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountXMASOccurancies(tt.grid)

			if result != tt.expected {
				t.Errorf("expected %d, got=%d", tt.expected, result)
			}
		})
	}
}
