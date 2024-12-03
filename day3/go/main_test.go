package main

import "testing"

func TestProcessInput(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"mul(2,3)do()mul(4,5)", 26},
		{"mul(1,2)don't()mul(2,3)", 2},
		{"mul(3,3)do()mul(2,2)don't()mul(1,1)", 13},
		{"don't()mul(10,10)", 0},
		{"mul(5,5)", 25},
	}

	for _, test := range tests {
		result := processInput(test.input)
		if result != test.expected {
			t.Errorf("processInput(%q), got = %d, want =  %d", test.input, result, test.expected)
		}
	}
}
