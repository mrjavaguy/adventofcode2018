package main

import "testing"

func TestDay06Part1(t *testing.T) {
	lines := []string{
		"1, 1",
		"1, 6",
		"8, 3",
		"3, 4",
		"5, 5",
		"8, 9",
	}
	result := Day06Part1(lines)
	if result != 17 {
		t.Errorf("Day06Part1 was incorrect, got: %d, want: %d.", result, 17)
	}
}
