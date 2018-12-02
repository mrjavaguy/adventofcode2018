package main

import (
	"adventofcode2018/input"
	"fmt"
	"strconv"
)

func convertStringArrayToIntArray(stringArray []string) []int {
	var intArray []int
	for _, str := range stringArray {
		x, _ := strconv.Atoi(str)
		intArray = append(intArray, x)
	}
	return intArray
}

func sum(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}

func day01Part1(input []int) int {
	return sum(input)
}

func day01Part2(input []int) int {
	found := false
	visited := make(map[int]bool)
	frequency := 0
	for !found {
		for i := range input {
			frequency += input[i]
			if _, ok := visited[frequency]; ok {
				found = true
				break
			}
			visited[frequency] = true
		}
	}

	return frequency
}

func main() {
	lines, _ := input.FileToLines("day01/input01.txt")
	day01input := convertStringArrayToIntArray(lines)
	part1 := day01Part1(day01input)
	part2 := day01Part2(day01input)
	fmt.Printf("Day 1 Part 1 %v, Part 2 %v", part1, part2)
}
