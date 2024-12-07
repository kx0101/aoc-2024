package main

import (
	"testing"
)

func TestConcatenate(t *testing.T) {
	tests := []struct {
		left, right, expected int64
	}{
		{12, 345, 12345},
		{0, 678, 678},
		{987, 65, 98765},
		{1, 2, 12},
	}

	for _, test := range tests {
		result := combineString(test.left, test.right)

		if result != test.expected {
			t.Errorf("combineString(%d, %d) = %d; want %d", test.left, test.right, result, test.expected)
		}
	}
}

func TestCanMatchTestValue(t *testing.T) {
	tests := []struct {
		target   int64
		numbers  []int64
		expected bool
	}{
		{190, []int64{10, 19}, true},        // 10 * 19 = 190
		{3267, []int64{81, 40, 27}, true},   // 81 + 40 * 27 = 3267
		{83, []int64{17, 5}, false},         // no combination
		{156, []int64{15, 6}, true},         // 15 || 6 = 156
		{7290, []int64{6, 8, 6, 15}, true},  // 6 * 8 || 6 * 15 = 7290
		{292, []int64{11, 6, 16, 20}, true}, // 11 + 6 * 16 + 20 = 292
		{9999, []int64{99, 99}, true},       // 99 || 99 = 9999
		{100000, []int64{100, 100}, false},   // no combination
	}

	for _, test := range tests {
		result := matchesValue(test.target, test.numbers)

		if result != test.expected {
			t.Errorf("matchesValue(%d, %v) = %v; want %v", test.target, test.numbers, result, test.expected)
		}
	}
}

func TestTryOperators(t *testing.T) {
	tests := []struct {
		numbers       []int64
		index         int
		currentResult int64
		target        int64
		expected      bool
	}{
		{[]int64{10, 19}, 1, 10, 190, true},        // 10 * 19 = 190
		{[]int64{15, 6}, 1, 15, 156, true},         // 15 || 6 = 156
		{[]int64{81, 40, 27}, 1, 81, 3267, true},   // 81 + 40 * 27 = 3267
		{[]int64{11, 6, 16, 20}, 1, 11, 292, true}, // 11 + 6 * 16 + 20 = 292
		{[]int64{99, 99}, 1, 99, 9999, true},       // 99 || 99 = 9999
		{[]int64{100, 100}, 1, 100, 1000, false},  // no combination
	}

	for _, test := range tests {
		result := tryOperators(test.numbers, test.index, test.currentResult, test.target)

		if result != test.expected {
			t.Errorf("tryOperators(%v, %d, %d, %d) = %v; want %v",
				test.numbers, test.index, test.currentResult, test.target, result, test.expected)
		}
	}
}
