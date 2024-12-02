package main

import "testing"

func TestIsSafe(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := isSafe(test.input)

		if result != test.expected {
			t.Errorf("isSafe(%v) = %v; want = %v", test.input, result, test.expected)
		}
	}
}

func TestCanBeSafeWithOneRemove(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, test := range tests {
		result := isSafe(test.input) || canBeSafeWithOneRemove(test.input)
		if result != test.expected {
			t.Errorf("canBeSafeWithOneRemove(%v) = %v; want = %v", test.input, result, test.expected)
		}
	}
}
