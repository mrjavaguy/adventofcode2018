package main

import (
	"adventofcode2018/input"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day19/input19.txt")
	program := ParseInput(lines)
	registers := make(Registers, 6)
	part1 := Day19Part1(program, registers)
	part2 := Day19Part2(program)
	fmt.Printf("Day 19 Part 1 %v, Part 2 %v", part1, part2)
}

func Day19Part1(program Program, registers Registers) int {
	registers = Run(program, registers)
	s := sum(registers[2])
	return s
}

func Day19Part2(program Program) int {
	registers := make(Registers, 6)
	registers[0] = 1
	registers = Run(program, registers)
	s := sum(registers[2])
	return s
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

func Run(program Program, registers Registers) Registers {
	for true {
		ip := registers[program.instructionPointer]
		if ip < 0 || ip > len(program.instructions) {
			break
		}

		if ip == 5 {
			break
		}

		next := program.instructions[ip]
		f := opCodes[next.op]
		registers = f(registers, next.a, next.b, next.c)

		registers[program.instructionPointer] = registers[program.instructionPointer] + 1
	}
	return registers
}

func sum(n int) int {
	s := 0
	r := int(math.Sqrt(float64(n)))
	for i := 2; i <= r; i++ {
		if n%i == 0 {
			if i == (n / i) {
				s += i
			} else {
				s += (i + n/i)
			}
		}
	}
	return s + 1 + n
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
