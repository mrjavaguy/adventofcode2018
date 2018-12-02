package main

import (
	"adventofcode2018/input"
	"fmt"
	"strings"
)

func checkLetterCount(input string) (bool, bool) {
	letter := make(map[byte]int)
	hasTwoChar := false
	hasThreeChar := false
	for i := 0; i < len(input); i++ {
		var idx = input[i]
		letter[idx] = letter[idx] + 1
	}

	for _, v := range letter {
		hasTwoChar = hasTwoChar || (v == 2)
		hasThreeChar = hasThreeChar || (v == 3)
	}

	return hasTwoChar, hasThreeChar
}

func calcCheckSum(input []string) int {
	count2 := 0
	count3 := 0
	for i := range input {
		hasTwoChar, hasThreeChar := checkLetterCount(input[i])
		if hasTwoChar {
			count2++
		}
		if hasThreeChar {
			count3++
		}
	}

	return count2 * count3
}

func stringDifference(x, y string) int {
	diff := 0
	for k := 0; k < len(x); k++ {
		if x[k] != y[k] {
			diff++
		}
	}
	return diff
}

func day02Part1(input []string) int {
	return calcCheckSum(input)
}

func day02Part2(input []string) string {
	var id1, id2 string
loop:
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			if stringDifference(input[i], input[j]) == 1 {
				id1, id2 = input[i], input[j]
				break loop
			}
		}
	}
	s := strings.Builder{}
	for i := 0; i < len(id1); i++ {
		if id1[i] == id2[i] {
			s.WriteByte(id1[i])
		}
	}

	return s.String()
}

func main() {
	lines, _ := input.FileToLines("day02/input02.txt")
	part1 := day02Part1(lines)
	part2 := day02Part2(lines)
	fmt.Printf("Day 2 Part 1 %v, Part 2 %v", part1, part2)
}
