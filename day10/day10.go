package main

import (
	"adventofcode2018/input"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, _ := input.FileToLines("day10/input10.txt")
	starList, average := parseLines(lines)
	solve(starList, average)
}

type star struct {
	pointX    int
	pointY    int
	velocityX int
	velocityY int
}

func parseLines(lines []string) (starList []star, average int) {
	re := regexp.MustCompile("position=<(.+), (.+)> velocity=<(.+), (.+)>")
	ave := make([]float64, 0)
	starList = make([]star, 0)
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		pX, _ := strconv.Atoi(strings.TrimSpace(matches[0][1]))
		pY, _ := strconv.Atoi(strings.TrimSpace(matches[0][2]))
		vX, _ := strconv.Atoi(strings.TrimSpace(matches[0][3]))
		vY, _ := strconv.Atoi(strings.TrimSpace(matches[0][4]))
		s := star{pointX: pX, pointY: pY, velocityX: vX, velocityY: vY}
		starList = append(starList, s)
		if vX == 0 {
			ave = append(ave, float64(pX))
		} else {
			ave = append(ave, math.Abs(float64(pX)/float64(vX)))
		}
		if vY == 0 {
			ave = append(ave, float64(pY))
		} else {
			ave = append(ave, math.Abs(float64(pY)/float64(vY)))
		}
	}

	sum := 0.0
	for _, a := range ave {
		sum += a
	}

	average = int(sum / float64(len(ave)))
	return
}

func solve(starList []star, average int) {
	savedList := make([]star, len(starList))
	for i, s := range starList {
		savedList[i] = s
	}

	for i, s := range starList {
		s.pointX += (s.velocityX * (average - 6))
		s.pointY += (s.velocityY * (average - 6))
		starList[i] = s
	}

	smallest := math.MaxInt64
	best := average - 11

	for start := average - 5; start < average+5; start++ {
		for i, s := range starList {
			s.pointX += s.velocityX
			s.pointY += s.velocityY
			starList[i] = s
		}

		sort.Slice(starList, func(i, j int) bool {
			return starList[i].pointX < starList[j].pointX
		})

		leftPoint := starList[0]
		rightPoint := starList[len(starList)-1]

		sort.Slice(starList, func(i, j int) bool {
			if starList[i].pointY == starList[j].pointY {
				return starList[i].pointX < starList[j].pointX
			}
			return starList[i].pointY < starList[j].pointY
		})

		topPoint := starList[0]
		bottomPoint := starList[len(starList)-1]

		area := int(math.Abs(float64(leftPoint.pointX-rightPoint.pointX) * math.Abs(float64(topPoint.pointY-bottomPoint.pointY))))
		if area < smallest {
			best = start
			smallest = area
		}
	}

	currentStar := 0
	for i, s := range savedList {
		s.pointX += (s.velocityX * best)
		s.pointY += (s.velocityY * best)
		savedList[i] = s
	}

	sort.Slice(savedList, func(i, j int) bool {
		return savedList[i].pointX < savedList[j].pointX
	})

	leftPoint := savedList[0]
	rightPoint := savedList[len(savedList)-1]

	sort.Slice(savedList, func(i, j int) bool {
		if savedList[i].pointY == savedList[j].pointY {
			return savedList[i].pointX < savedList[j].pointX
		}
		return savedList[i].pointY < savedList[j].pointY
	})

	topPoint := savedList[0]
	bottomPoint := savedList[len(savedList)-1]

	// Part 1
	fmt.Println("Part 1 ")
	for y := topPoint.pointY; y <= bottomPoint.pointY; y++ {
		for x := leftPoint.pointX; x <= rightPoint.pointX; x++ {
			s := savedList[currentStar]
			if s.pointX == x && s.pointY == y {
				fmt.Print("#")
				for s.pointX == x && s.pointY == y {
					currentStar++
					if currentStar == len(savedList) {
						currentStar--
						break
					}
					s = savedList[currentStar]
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	// Part 2
	fmt.Println("Part 2 ", best)
}
