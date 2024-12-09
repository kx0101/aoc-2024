package main

import (
	"reflect"
	"testing"
)

func TestParseDisk(t *testing.T) {
	tests := []struct {
		input    string
		expected []rune
	}{
		{"2333133121414131402", []rune("00...111...2...333.44.5555.6666.777.888899")},
		{"12345", []rune("0..111....22222")},
	}

	for _, test := range tests {
		actual := parseDisk(test.input)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("parseDisk(%q) = %q; want=%q", test.input, string(actual), string(test.expected))
		}
	}
}

func TestMoveBlocks(t *testing.T) {
	tests := []struct {
		input    []rune
		expected []rune
	}{
		{
			[]rune("00...111...2...333.44.5555.6666.777.888899"),
			[]rune("00992111777.44.333....5555.6666.....8888.."),
		},
		{
			[]rune("0..111....22222"),
			[]rune("0..111....22222"),
		},
		{
			[]rune("000000000.000000000"),
			[]rune("000000000.000000000"),
		},
	}

	for _, test := range tests {
		actual := make([]rune, len(test.input))

		copy(actual, test.input)
		moveBlocks(actual)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("moveBlocks(%q) = %q; want=%q", string(test.input), string(actual), string(test.expected))
		}
	}
}

func TestCalculateChecksum(t *testing.T) {
	tests := []struct {
		input    []rune
		expected int64
	}{
		{[]rune("00992111777.44.333....5555.6666.....8888.."), 2858},
		{[]rune("0..111....22222"), 132},
	}

	for _, test := range tests {
		actual := calculateChecksum(test.input)

		if actual != test.expected {
			t.Errorf("calculateChecksum(%q) = %d; want=%d", string(test.input), actual, test.expected)
		}
	}
}
