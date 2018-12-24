package main

import "testing"

func TestDay22Part1(t *testing.T) {
	part1 := Day22Part1(510, complex(10, 10))
	expected := 114
	if part1 != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", part1, expected)
	}
}

func TestDay22Part2(t *testing.T) {
	part1 := Day22Part2(510, complex(10, 10))
	expected := 45
	if part1 != expected {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", part1, expected)
	}
}
