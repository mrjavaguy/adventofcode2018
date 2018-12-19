package main

import "testing"

var lines = []string{
	"x=495, y=2..7",
	"y=7, x=495..501",
	"x=501, y=3..7",
	"x=498, y=2..4",
	"x=506, y=1..2",
	"x=498, y=10..13",
	"x=504, y=10..13",
	"y=13, x=498..504",
}

func TestDay17Part1(t *testing.T) {
	groundMap := ParseInput(lines)
	water, flow := Solve(groundMap)
	if water != 29 {
		t.Errorf("Solve was incorrect for water, got: %d, want: %d.", water, 29)
	}
	if flow != 28 {
		t.Errorf("Solve was incorrect for flow, got: %d, want: %d.", water, 28)
	}
	PrintMap(groundMap)
}
