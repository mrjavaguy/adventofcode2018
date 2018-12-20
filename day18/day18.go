package main

import (
	"adventofcode2018/input"
	"fmt"
)

func main() {
	lines, _ := input.FileToLines("day18/input18.txt")
	area := ParseInput(lines)
	part1 := Day18Part1(area)
	area = ParseInput(lines)
	part2 := Day18Part2(area)
	fmt.Printf("Day 18 Part 1 %v, Part 2 %v", part1, part2)
}

type collectionArea map[complex64]rune

func ParseInput(lines []string) (area collectionArea) {
	area = collectionArea{}
	for y, l := range lines {
		for x, c := range l {
			area[makePoint(x, y)] = c
		}
	}
	return
}

func makePoint(x, y int) complex64 {
	return complex(float32(x), float32(y))
}

var offsets = []complex64{
	complex(-1, -1),
	complex(0, -1),
	complex(1, -1),
	complex(-1, 0),
	complex(1, 0),
	complex(-1, 1),
	complex(0, 1),
	complex(1, 1),
}

func getNeighbors(area collectionArea, p complex64) (trees, lumberyards, open int) {
	neighbors := []rune{}
	for _, x := range offsets {
		if c, ok := area[p+x]; ok {
			neighbors = append(neighbors, c)
		}
	}
	return neighborTypeCount(neighbors)
}

func neighborTypeCount(neighbors []rune) (trees, lumberyards, open int) {
	for _, c := range neighbors {
		switch c {
		case '|':
			trees++
		case '#':
			lumberyards++
		case '.':
			open++
		default:
			panic("invalid rune found " + string(c))
		}
	}
	return
}

func areaTypeCount(area collectionArea) (trees, lumberyards, open int) {
	for _, c := range area {
		switch c {
		case '|':
			trees++
		case '#':
			lumberyards++
		case '.':
			open++
		default:
			panic("invalid rune found " + string(c))
		}
	}
	return
}

func gameOfLife(area collectionArea) (newArea collectionArea) {
	newArea = collectionArea{}
	for p, c := range area {
		newArea[p] = c
		trees, lumberyards, _ := getNeighbors(area, p)
		switch c {
		case '.':
			if trees > 2 {
				newArea[p] = '|'
			}
		case '|':
			if lumberyards > 2 {
				newArea[p] = '#'
			}
		case '#':
			if trees == 0 || lumberyards == 0 {
				newArea[p] = '.'
			}
		}
		//fmt.Println(p, string(c), trees, lumberyards, string(newArea[p]))
	}
	return
}

func printMap(area collectionArea) {
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			fmt.Print(string(area[makePoint(x, y)]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Day18Part1(area collectionArea) int {
	//fmt.Println("Initial state:")
	//printMap(area)
	for t := 1; t <= 10; t++ {
		area = gameOfLife(area)
		//fmt.Println("After", t, "minute(s):")
		//printMap(area)
	}
	trees, lumberyards, _ := areaTypeCount(area)
	return trees * lumberyards
}

func compare(area1, area2 collectionArea) bool {
	for i := range area1 {
		if area1[i] != area2[i] {
			return false
		}
	}
	return true
}

func contains(memo []collectionArea, area collectionArea) (bool, int) {
	for i, a := range memo {
		if compare(a, area) {
			return true, i
		}
	}
	return false, -1
}

func Day18Part2(area collectionArea) int {
	iteration := 0
	cycleLength := 0
	idx := 0
	ok := false
	memo := []collectionArea{}
	for true {
		area = gameOfLife(area)
		if ok, idx = contains(memo, area); ok {
			cycleLength = len(memo) - idx
			break
		}
		memo = append(memo, area)
		iteration++
	}

	left := idx + (1000000000-idx)%cycleLength - 1

	fmt.Println(cycleLength, left, iteration)

	trees, lumberyards, _ := areaTypeCount(memo[left])
	return trees * lumberyards
}
