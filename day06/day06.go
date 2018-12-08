package main

import (
	"adventofcode2018/input"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day06/input06.txt")
	part1 := Day06Part1(lines)
	part2 := Day06Part2(lines)
	fmt.Printf("Day 6 Part 1 %v, Part 2 %v", part1, part2)
}

func Day06Part1(lines []string) int {
	points := parseLines(lines)
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	leftPoint := points[0]
	rightPoint := points[len(points)-1]

	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})

	topPoint := points[0]
	bottomPoint := points[len(points)-1]

	count := make(map[Point]int)
	invalid := make(map[Point]bool)

	for x := leftPoint.X - 1; x < rightPoint.X+1; x++ {
		for y := topPoint.Y - 1; y < bottomPoint.Y+1; y++ {
			d := math.MaxInt32
			least := new(Point)
			for _, p := range points {
				c := calcDistance(p, x, y)
				if c == d {
					least = nil
				}
				if c < d {
					if least == nil {
						least = new(Point)
					}
					*least = p
					d = c
				}
			}
			if least != nil {
				count[*least]++
				if (x == leftPoint.X-1) || (x == rightPoint.X) || (y == topPoint.Y-1) || (y == bottomPoint.Y) {
					invalid[*least] = true
				}
			}
		}
	}

	for k, b := range invalid {
		if b {
			delete(count, k)
		}
	}

	max := -1

	for _, i := range count {
		if i > max {
			max = i
		}
	}

	return max
}

func Day06Part2(lines []string) int {
	points := parseLines(lines)
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})

	leftPoint := points[0]
	rightPoint := points[len(points)-1]

	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})

	topPoint := points[0]
	bottomPoint := points[len(points)-1]

	count := make(map[Point]int)
	region := 0

	for x := leftPoint.X; x < rightPoint.X; x++ {
		for y := topPoint.Y; y < bottomPoint.Y; y++ {
			cp := Point{X: x, Y: y}
			for _, p := range points {
				c := calcDistance(p, x, y)
				count[cp] += c
			}
			if count[cp] < 10000 {
				region++
			}
		}
	}
	return region
}

type Point struct {
	X int
	Y int
}

func calcDistance(p Point, x int, y int) int {
	return int(math.Abs(float64(p.X-x)) + math.Abs(float64(p.Y-y)))
}

func parseLines(lines []string) []Point {
	points := make([]Point, len(lines))
	re := regexp.MustCompile("(\\d+), (\\d+)")

	for i, s := range lines {
		matches := re.FindAllStringSubmatch(s, -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		points[i] = Point{X: x, Y: y}
	}

	return points
}
