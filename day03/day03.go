package main

import (
	"adventofcode2018/input"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines, _ := input.FileToLines("day03/input03.txt")
	rects := createRectangles(lines)
	board := createFabricBoard(rects)
	part1 := day03Part1(board)
	part2 := day03Part2(rects, board)
	fmt.Printf("Day 3 Part 1 %v, Part 2 %v", part1, part2)
}

func day03Part1(board map[Vertex][]string) int {
	area := 0
	for _, v := range board {
		if len(v) > 1 {
			area++
		}
	}
	return area
}

func day03Part2(input []Rectangle, board map[Vertex][]string) string {
outer:
	for _, rect := range input {
		for x := rect.TopLeft.X; x < rect.TopLeft.X+rect.WidthHeight.X; x++ {
			for y := rect.TopLeft.Y; y < rect.TopLeft.Y+rect.WidthHeight.Y; y++ {
				if len(board[Vertex{x, y}]) > 1 {
					continue outer
				}
			}
		}

		return rect.ID
	}
	return ""
}

type Vertex struct {
	X, Y int
}

type Rectangle struct {
	ID          string
	TopLeft     Vertex
	WidthHeight Vertex
}

func createRectangles(input []string) []Rectangle {
	var rects []Rectangle
	for i := range input {
		rects = append(rects, parseRectangle(input[i]))
	}
	return rects
}

func parseRectangle(input string) Rectangle {
	re := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")
	matches := re.FindAllStringSubmatch(input, -1)
	id := matches[0][1]
	x, _ := strconv.Atoi(matches[0][2])
	y, _ := strconv.Atoi(matches[0][3])
	w, _ := strconv.Atoi(matches[0][4])
	h, _ := strconv.Atoi(matches[0][5])
	rect := Rectangle{
		ID:          id,
		TopLeft:     Vertex{X: x, Y: y},
		WidthHeight: Vertex{X: w, Y: h}}
	return rect
}

func createFabricBoard(rects []Rectangle) map[Vertex][]string {
	fabric := make(map[Vertex][]string)
	for _, rect := range rects {
		for x := rect.TopLeft.X; x < rect.TopLeft.X+rect.WidthHeight.X; x++ {
			for y := rect.TopLeft.Y; y < rect.TopLeft.Y+rect.WidthHeight.Y; y++ {
				fabric[Vertex{x, y}] = append(fabric[Vertex{x, y}], rect.ID)
			}
		}
	}
	return fabric
}
