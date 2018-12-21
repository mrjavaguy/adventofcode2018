package main

import "testing"

var lines = []string{
	"#ip 0",
	"seti 5 0 1",
	"seti 6 0 2",
	"addi 0 1 0",
	"addr 1 2 3",
	"setr 1 0 0",
	"seti 8 0 4",
	"seti 9 0 5",
}

func TestDay19Part1(t *testing.T) {
	program := ParseInput(lines)
	registers := make(Registers, 6)
	part1 := Day19Part1(program, registers)
	expected := 7
	if part1 != expected {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", part1, expected)
	}
}
