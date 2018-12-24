package main

import "testing"

var lines = []string{
	"pos=<0,0,0>, r=4",
	"pos=<1,0,0>, r=1",
	"pos=<4,0,0>, r=3",
	"pos=<0,2,0>, r=1",
	"pos=<0,5,0>, r=3",
	"pos=<0,0,3>, r=1",
	"pos=<1,1,1>, r=1",
	"pos=<1,1,2>, r=1",
	"pos=<1,3,1>, r=1",
}

func TestDay23Part1(t *testing.T) {
	swarm, masterBot := ParseInput(lines)

	part1 := Day23Part1(swarm, masterBot)
	expected := 7
	if part1 != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", part1, expected)
	}
}

var lines2 = []string{
	"pos=<10,12,12>, r=2",
	"pos=<12,14,12>, r=2",
	"pos=<16,12,12>, r=4",
	"pos=<14,14,14>, r=6",
	"pos=<50,50,50>, r=200",
	"pos=<10,10,10>, r=5",
}

func TestDay23Part2(t *testing.T) {
	swarm, _ := ParseInput(lines2)

	part2 := Day23Part2(swarm)
	expected := 36
	if part2 != expected {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", part2, expected)
	}
}
