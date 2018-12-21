package main

import (
	"adventofcode2018/input"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day21/input21.txt")
	program := ParseInput(lines)
	part1 := Day21Part1(program)
	part2 := Day21Part2(program)
	fmt.Printf("Day 21 Part 1 %v, Part 2 %v", part1, part2)
}

func Day21Part1(program Program) int {
	registers := make(Registers, 6)
	return Run(program, registers, true)
}

func Day21Part2(program Program) int {
	registers := make(Registers, 6)
	return Run(program, registers, false)
}

type ProgramLine struct {
	op string
	a  int
	b  int
	c  int
}

type Program struct {
	instructionPointer int
	instructions       []ProgramLine
}

type Registers []int

func ParseInput(lines []string) (program Program) {
	program = Program{
		instructions: []ProgramLine{},
	}
	re := regexp.MustCompile("(\\d+)")
	matches := re.FindAllStringSubmatch(lines[0], -1)
	program.instructionPointer, _ = strconv.Atoi(matches[0][0])
	re = regexp.MustCompile("(\\w+) (\\d+) (\\d+) (\\d+)")
	for i, l := range lines {
		if i == 0 {
			continue
		}
		matches = re.FindAllStringSubmatch(l, -1)
		A, _ := strconv.Atoi(matches[0][2])
		B, _ := strconv.Atoi(matches[0][3])
		C, _ := strconv.Atoi(matches[0][4])
		program.instructions = append(program.instructions, ProgramLine{
			op: matches[0][1],
			a:  A,
			b:  B,
			c:  C,
		})
	}

	return
}

func Run(program Program, registers Registers, exitQuickly bool) int {
	values := map[int]struct{}{}
	last := -1
	for true {
		ip := registers[program.instructionPointer]
		if ip < 0 || ip > len(program.instructions) {
			break
		}

		if ip == 28 {
			if exitQuickly {
				return registers[3]
			} else {
				_, ok := values[registers[3]]
				if ok {
					return last
				}
				values[registers[3]] = struct{}{}
				last = registers[3]
			}
		}

		next := program.instructions[ip]
		f := opCodes[next.op]
		registers = f(registers, next.a, next.b, next.c)

		registers[program.instructionPointer] = registers[program.instructionPointer] + 1
	}
	return -1
}

func addr(registers []int, A, B, C int) []int {
	registers[C] = registers[A] + registers[B]
	return registers
}

func addi(registers []int, A, B, C int) []int {
	registers[C] = registers[A] + B
	return registers
}

func mulr(registers []int, A, B, C int) []int {
	registers[C] = registers[A] * registers[B]
	return registers
}

func muli(registers []int, A, B, C int) []int {
	registers[C] = registers[A] * B
	return registers
}

func banr(registers []int, A, B, C int) []int {
	registers[C] = registers[A] & registers[B]
	return registers
}

func bani(registers []int, A, B, C int) []int {
	registers[C] = registers[A] & B
	return registers
}

func borr(registers []int, A, B, C int) []int {
	registers[C] = registers[A] | registers[B]
	return registers
}

func bori(registers []int, A, B, C int) []int {
	registers[C] = registers[A] | B
	return registers
}

func setr(registers []int, A, B, C int) []int {
	registers[C] = registers[A]
	return registers
}

func seti(registers []int, A, B, C int) []int {
	registers[C] = A
	return registers
}

func gtir(registers []int, A, B, C int) []int {
	if A > registers[B] {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

func gtri(registers []int, A, B, C int) []int {
	if registers[A] > B {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

func gtrr(registers []int, A, B, C int) []int {
	if registers[A] > registers[B] {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

func eqir(registers []int, A, B, C int) []int {
	if A == registers[B] {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

func eqri(registers []int, A, B, C int) []int {
	if registers[A] == B {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

func eqrr(registers []int, A, B, C int) []int {
	if registers[A] == registers[B] {
		registers[C] = 1
	} else {
		registers[C] = 0
	}
	return registers
}

type opCode func([]int, int, int, int) []int

type OpCodes map[string]opCode

var opCodes = OpCodes{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}
