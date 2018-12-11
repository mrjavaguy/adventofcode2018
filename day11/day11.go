package main

import (
	"fmt"
	"math"
)

func main() {

	serialNumber := 4151
	grid := calculatePowerLevel(serialNumber)
	x, y, _ := findMostFuel(grid, 3)
	fmt.Println("Day 11 Part 1", x, y)
	x2, y2, s := partialsums(serialNumber)
	fmt.Println("Day 11 Part 2", x2, y2, s)
}

type point struct {
	x int
	y int
}

func calculatePowerLevel(serialNumber int) (grid map[point]int) {
	grid = map[point]int{}
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			grid[point{x: x, y: y}] = powerLevel(x, y, serialNumber)
		}
	}
	return
}

func powerLevel(x, y, serialNumber int) int {
	rackID := x + 10
	power := rackID * y
	power += serialNumber
	power *= rackID
	power /= 100
	power = power % 10
	return power - 5
}

func findMostFuel(grid map[point]int, size int) (x, y, maxFuel int) {
	maxFuel = -1
	x = -1
	y = -1
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			fuelLevel := 0
			for l := 0; l < size; l++ {
				for k := 0; k < size; k++ {
					fuelLevel += grid[point{x: i + l, y: j + k}]
				}
			}
			if fuelLevel > maxFuel {
				maxFuel = fuelLevel
				x = i
				y = j
			}
		}
	}
	return
}

func partialsums(serialNumber int) (xValue, yValue, size int) {
	grid := [301][301]int{}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			grid[y][x] = powerLevel(x, y, serialNumber) + grid[y-1][x] + grid[y][x-1] - grid[y-1][x-1]
		}
	}

	xValue = math.MinInt32
	yValue = math.MinInt32
	size = math.MinInt32
	best := math.MinInt32
	for s := 1; s <= 300; s++ {
		for y := s; y <= 300; y++ {
			for x := s; x <= 300; x++ {
				total := grid[y][x] - grid[y-s][x] - grid[y][x-s] + grid[y-s][x-s]
				if total > best {
					best = total
					xValue = x
					yValue = y
					size = s
				}
			}
		}
	}
	xValue = xValue - size + 1
	yValue = yValue - size + 1
	return
}
