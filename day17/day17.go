package main

import (
	"adventofcode2018/input"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day17/input17.txt")

	groundMap := ParseInput(lines)
	part1, part2 := Solve(groundMap)
	fmt.Printf("Day 17 Part 1 %v, Part 2 %v", part1+part2, part2)
}

type GroundMap struct {
	Map      [2000][2000]rune
	MaxDepth int
	MaxWidth int
	MinDepth int
	MinWidth int
}

func ParseInput(lines []string) (groundMap *GroundMap) {
	groundMap = &GroundMap{}
	maxX, maxY := 0, 0
	minX, minY := 2000, 2000
	re := regexp.MustCompile("(.)=(\\d+), (.)=(\\d+)..(\\d+)")
	for _, l := range lines {
		matches := re.FindAllStringSubmatch(l, -1)
		startLetter := matches[0][1]
		startIdx, _ := strconv.Atoi(matches[0][2])
		endIdx1, _ := strconv.Atoi(matches[0][4])
		endIdx2, _ := strconv.Atoi(matches[0][5])
		if startLetter == "x" {
			if startIdx > maxX {
				maxX = startIdx
			}
			if startIdx < minX {
				minX = startIdx
			}
			for i := endIdx1; i <= endIdx2; i++ {
				if i < minY {
					minY = i
				}
				if i > maxY {
					maxY = i
				}
				groundMap.Map[startIdx][i] = '#'
			}

		} else {
			if startIdx > maxY {
				maxY = startIdx
			}
			if startIdx < minY {
				minY = startIdx
			}
			for i := endIdx1; i <= endIdx2; i++ {
				if i < minX {
					minX = i
				}
				if i > maxX {
					maxX = i
				}
				groundMap.Map[i][startIdx] = '#'
			}
		}
	}
	groundMap.MaxDepth = maxY
	groundMap.MaxWidth = maxX
	groundMap.MinWidth = minX
	groundMap.MinDepth = minY
	return
}

func open(x, y int, groundMap *GroundMap) bool {
	return groundMap.Map[x][y] == 0 || groundMap.Map[x][y] == '|'
}

func fill(x, y int, groundMap *GroundMap) {
	if y > groundMap.MaxDepth {
		return
	} else if !open(x, y, groundMap) {
		return
	}
	if !open(x, y+1, groundMap) {
		leftX := x
		for open(leftX, y, groundMap) && !open(leftX, y+1, groundMap) {
			groundMap.Map[leftX][y] = '|'
			leftX--
		}
		rightX := x + 1
		for open(rightX, y, groundMap) && !open(rightX, y+1, groundMap) {
			groundMap.Map[rightX][y] = '|'
			rightX++
		}
		if open(leftX, y+1, groundMap) || open(rightX, y+1, groundMap) {
			fill(leftX, y, groundMap)
			fill(rightX, y, groundMap)
		} else if groundMap.Map[leftX][y] == '#' && groundMap.Map[rightX][y] == '#' {
			for x2 := leftX + 1; x2 < rightX; x2++ {
				groundMap.Map[x2][y] = '~'
			}
		}
	} else if groundMap.Map[x][y] == 0 {
		groundMap.Map[x][y] = '|'
		fill(x, y+1, groundMap)
		if groundMap.Map[x][y+1] == '~' {
			fill(x, y, groundMap)
		}
	}
}

func PrintMap(groundMap *GroundMap) {
	for y := 0; y <= groundMap.MaxDepth+1; y++ {
		for x := groundMap.MinWidth - 1; x < groundMap.MaxWidth+1; x++ {
			fmt.Print(string(groundMap.Map[x][y]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func Solve(groundMap *GroundMap) (int, int) {
	fill(500, 0, groundMap)

	water, touched := 0, 0
	for x := groundMap.MinWidth - 1; x <= groundMap.MaxWidth+1; x++ {
		for y := groundMap.MinDepth; y <= groundMap.MaxDepth; y++ {
			if groundMap.Map[x][y] == '|' {
				touched++
			} else if groundMap.Map[x][y] == '~' {
				water++
			}
		}
	}
	return water, touched
}
