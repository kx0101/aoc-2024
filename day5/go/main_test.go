package main

import (
	"testing"
)

func TestParseUpdate(t *testing.T) {
	input := "75,47,61,53,29"
	expected := []int{75, 47, 61, 53, 29}

	result := parseUpdateToInts(input)

	if !equalSlices(result, expected) {
		t.Errorf("expected: %v, got=%v\n", expected, result)
	}
}

func TestIsValidOrder(t *testing.T) {
	pages := []int{75, 47, 61, 53, 29}
	rules := [][2]int{
		{75, 47},
		{47, 61},
		{61, 53},
		{53, 29},
	}

	if !isValidOrder(pages, rules) {
		t.Errorf("expected to be valid, but it's not")
	}

	invalidPages := []int{53, 47, 61, 75, 29}
	if isValidOrder(invalidPages, rules) {
		t.Errorf("expected to be invalid, but it's not")
	}
}

func TestTopologicalSort(t *testing.T) {
	pages := []int{75, 47, 61, 53, 29}
	rules := [][2]int{
		{75, 47},
		{47, 61},
		{61, 53},
		{53, 29},
	}
	expected := []int{75, 47, 61, 53, 29}
	result := topologicalSort(pages, rules)

	if !equalSlices(result, expected) {
		t.Errorf("expected: %v, got=%v", expected, result)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
