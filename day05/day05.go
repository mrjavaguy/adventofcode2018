package main

import (
	"adventofcode2018/input"
	"fmt"
	"strings"
)

func main() {

	line, _ := input.FileToLine("day05/input05.txt")
	//line := "dabAcCaCBAcCcaDA"

	part1 := day05Part1(line)
	part2 := day05Part2(line)
	fmt.Printf("Day 1 Part 1 %v, Part 2 %v", part1, part2)
}

func day05Part1(line string) int {
	return colapse(line)
}

func day05Part2(line string) int {
	min := len(line)
	for i := rune('a'); i <= rune('z'); i++ {
		toColapse := strings.Replace(line, string(i), "", -1)
		toColapse = strings.Replace(toColapse, string(swapRune(i)), "", -1)
		calc := colapse(toColapse)
		if min > calc {
			min = calc
		}
	}

	return min
}

func colapse(line string) int {
	runes := []rune(line)
	runes = append(runes, rune(0))
	i := 0
	for i < len(runes)-1 {
		if runes[i] == swapRune(runes[i+1]) {
			runes = append(runes[:i], runes[i+2:]...)
			i--
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}

	return len(runes) - 1
}

func swapRune(r rune) rune {
	switch {
	case 97 <= r && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}
