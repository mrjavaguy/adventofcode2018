package main

import "testing"

var lines = []string{
	".#.#...|#.",
	".....#|##|",
	".|..|...#.",
	"..|#.....#",
	"#.#|||#|#|",
	"...#.||...",
	".|....|...",
	"||...#|.#|",
	"|.||||..|.",
	"...#.|..|.",
}

func TestDay19Part1(t *testing.T) {
	area := ParseInput(lines)
	part1 := Day18Part1(area)
	expected := 1147
	if part1 != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", part1, expected)
	}
}
