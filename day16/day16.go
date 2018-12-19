package main

import (
	"adventofcode2018/input"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day16/input16.txt")
	samples := parsePart1(lines)
	part1 := Day16Part1(samples)
	part2 := Day16Part2(samples)
	fmt.Printf("Day 16 Part 1 %v, Part 2 %v", part1, part2)
}

func Day16Part1(samples []*Sample) int {
	for _, s := range samples {
		for _, f := range opCodes {
			start := []int{
				s.StartState[0],
				s.StartState[1],
				s.StartState[2],
				s.StartState[3],
			}
			end := f(start, s.Code[1], s.Code[2], s.Code[3])
			if end[0] == s.EndState[0] && end[1] == s.EndState[1] && end[2] == s.EndState[2] && end[3] == s.EndState[3] {
				s.Possible = append(s.Possible, f)
			}
		}
	}
	i := 0
	for _, s := range samples {
		if len(s.Possible) > 2 {
			i++
		}
	}

	return i
}

func Day16Part2(samples []*Sample) int {
	opsByCode := OpsByCode{}
	for len(opsByCode) < 16 {
		for _, s1 := range samples {
			for _, s2 := range samples {
				if s1.Code[0] == s2.Code[0] {
					p := intersect(s1.Possible, s2.Possible)
					s1.Possible = p
					s2.Possible = p
					if len(s1.Possible) == 1 {
						opsByCode[s1.Code[0]] = s1.Possible[0]
						sf1 := reflect.ValueOf(s1.Possible[0])
						for _, s3 := range samples {
							if s1.Code[0] != s3.Code[0] {
								set := make([]opCode, 0)
								for _, p := range s3.Possible {
									sf2 := reflect.ValueOf(p)
									if sf1.Pointer() != sf2.Pointer() {
										set = append(set, p)
									}
								}
								s3.Possible = set
							}
						}
					}
				}
			}
		}
	}

	lines, _ := input.FileToLines("day16/testprogram.txt")
	fmt.Println(len(lines))
	regs := []int{0, 0, 0, 0}
	re := regexp.MustCompile("(\\d+)")
	for _, l := range lines {
		matches := re.FindAllStringSubmatch(l, -1)
		op, _ := strconv.Atoi(matches[0][0])
		a, _ := strconv.Atoi(matches[1][0])
		b, _ := strconv.Atoi(matches[2][0])
		c, _ := strconv.Atoi(matches[3][0])
		opsByCode[op](regs, a, b, c)
	}

	return regs[0]
}

func intersect(a []opCode, b []opCode) []opCode {
	set := make([]opCode, 0)

	for _, c1 := range a {
		for _, c2 := range b {
			sf1 := reflect.ValueOf(c1)
			sf2 := reflect.ValueOf(c2)
			if sf1.Pointer() == sf2.Pointer() {
				set = append(set, c1)
			}
		}
	}

	return set
}

type Sample struct {
	StartState []int
	EndState   []int
	Code       []int
	Possible   []opCode
}

func NewSample(lines []string) Sample {
	sample := Sample{
		StartState: make([]int, 4),
		EndState:   make([]int, 4),
		Code:       make([]int, 4),
		Possible:   make([]opCode, 0),
	}
	re := regexp.MustCompile("(\\d+)")
	matches := re.FindAllStringSubmatch(lines[0], -1)
	for x := 0; x < 4; x++ {
		sample.StartState[x], _ = strconv.Atoi(matches[x][0])
	}
	matches = re.FindAllStringSubmatch(lines[1], -1)
	for x := 0; x < 4; x++ {
		sample.Code[x], _ = strconv.Atoi(matches[x][0])
	}
	matches = re.FindAllStringSubmatch(lines[2], -1)
	for x := 0; x < 4; x++ {
		sample.EndState[x], _ = strconv.Atoi(matches[x][0])
	}
	return sample
}

func parsePart1(lines []string) []*Sample {
	samples := []*Sample{}
	sampleLine := make([]string, 4)
	for i, l := range lines {
		if i%4 == 3 {
			s := NewSample(sampleLine)
			samples = append(samples, &s)
		}
		sampleLine[i%4] = l
	}

	return samples
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

type OpsByCode map[int]opCode

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
